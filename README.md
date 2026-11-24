# lisette-collections

Clojure-style persistent collections and transducers for Lisette (Rust-like
syntax, Go runtime). Generic and monomorphized — no Go `any` boxing in hot
loops. Extracted from the N3 exam rapid practice app, where the HAMT card index
cut `/api/plan` latency from ~916 ms to ~109 ms.

## Architecture — three layers

```
Layer 3  collection API: map · filter · reduce · reduce-kv · conj · assoc ·
         dissoc · get · seq · nth · into · transduce · union/intersection/difference
Layer 2  traversal/transformation: Transducers · Reducers · Collection/KvCollection
         interfaces · iterators (indexed trie, linked list, map entry walk)
Layer 1  persistent storage: 32-way vector trie + compact tail · HAMT ·
         set-over-map · cons list · two-list queue · array-backed leaves
```

Files in `src/collections/`:

| File | Contents |
|---|---|
| `transducer.lis` | Reduced/RState, Stage/Pipeline, map/filter/take/drop/compose, one-loop run/fold/find |
| `persistent.lis` | cons List, 32-way Vector trie, HAMT HashMap (small-map ≤8), HashSet wrapper, two-list Queue |
| `smap.lis` | string-keyed HAMT (FNV-1a, 5-bit chunks, path-copy assoc), reduce-kv, transduce_keys |
| `protocol.lis` | Lisette interfaces `Collection<T>`, `KvCollection<K,V>`; per-structure traversal; transduce/into |
| `ops.lis` | set algebra, map merge/transform/filter, slice equality |

`model/` in the parent project (`n3examrapid`) contains the Quint executable
specification (`transducerCollections.qnt`) with 9 laws + 4 negative controls.

## Law/property tests

`lis test` → 21 passing: fusion (fused == materialized), category laws,
persistence (originals unchanged), trie nth across the 32-boundary, HAMT
assoc/dissoc, set-over-map, two-list queue, one transducer over five
collections, set algebra, map transforms.

## Usage

```sh
lis check
lis test
```

No external dependencies beyond the Go standard library (`slices`, `strings`).

## Future plans

1. **Persistent data structures using transducer abstractions and Lisette's
   type system.** Push more of the API through `Collection<T>` /
   `KvCollection<K,V>` interfaces; specialize traversal per representation
   (indexed/chunk-aware for vectors, linked for lists, direct entry walk for
   maps) while keeping one transformation protocol.
2. **Transient / mutable-at-first then freeze collections.** Go slices are
   excellent transients and dense backing stores: build with in-place mutable
   buffers, then freeze into a persistent trie/HAMT with structural sharing.
   Bulk-load sorted leaves (B+tree lesson: detect ordered input, keep first
   key + last value, merge batches; plain conj for ≤8 entries).
3. **Compatibility with Go GC pressure.** Investigate which paradigms fit Go's
   memory model: per-node allocation vs contiguous leaves, pointer-heavy
   tries vs chunked arrays, arena/bump allocation for builders, escape
   analysis, and measurement with pprof/massif. Sorted map/set should use a
   persistent B+tree (16-entry leaves, 32-child inner nodes) rather than a
   red-black tree — see <https://yogthos.net/posts/2026-11-23-b+trees.html>.
   Always measure small, medium, and large sizes; memory layout dominates.

## Non-goals

Runtime laziness/thunks/coinductive streams, zero *total* allocation (result
structures are allocated), and sorted collections (API reserved, implementation
deferred to the B+tree plan).
