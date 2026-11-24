# lisette-persistent-collection

Clojure-style persistent collections and transducers for Lisette. The library is generic and monomorphized, so hot loops avoid `any` boxing. In an Anki-style web app, using the String-keyed persistent HAMT for card index reduced `/api/plan` latency from roughly 916 ms to 109 ms.

## Architecture

The library has three layers:

```
Layer 3  Collection API — map, filter, reduce, reduce-kv, conj, assoc,
         dissoc, get, seq, nth, into, transduce, and set operations
Layer 2  Traversal and transformation — transducers, reducers, the
         Collection/KvCollection interfaces, and iterators (indexed trie,
         linked list, map-entry walk)
Layer 1  Persistent storage — a 32-way vector trie with a compact tail,
         HAMT, set-over-map, cons list, two-list queue, and array-backed leaves
```

Files in `src/collections/`:

| File | Contents |
|---|---|
| `transducer.lis` | `Reduced`/`RState`, `Stage`/`Pipeline`, map/filter/take/drop/compose, and one-loop run/fold/find |
| `persistent.lis` | cons list, 32-way vector trie, HAMT hash map (small maps of up to eight entries), `HashSet` wrapper, two-list queue |
| `smap.lis` | string-keyed HAMT (FNV-1a, 5-bit chunks, path-copying `assoc`), `reduce-kv`, `transduce_keys` |
| `protocol.lis` | Lisette interfaces `Collection<T>` and `KvCollection<K,V>`, per-structure traversal, `transduce`/`into` |
| `ops.lis` | set algebra, map merge/transform/filter, slice equality |


## Tests

`lis test` runs 21 property tests covering:

- Transducer fusion (fused results match materialized ones)
- Category laws
- Persistence (source collections remain unchanged)
- Vector indexing across the 32-element boundary
- HAMT `assoc` and `dissoc`
- set-over-map
- The two-list queue
- One transducer applied across five collection types
- Set algebra
- Map transforms

## Usage

```sh
lis check
lis test
```

The library has no dependencies beyond the Go standard library (`slices` and `strings`).

## Future work

1. **Deepen the collection abstraction.** Move more of the API behind the `Collection<T>` and `KvCollection<K,V>` interfaces, specializing traversal per representation (indexed and chunk-aware for vectors, linked for lists, direct entry walks for maps) while keeping a single transformation protocol.
2. **Add transient builders.** Use Go slices as mutable buffers during construction, then freeze them into a persistent trie or HAMT with structural sharing. Bulk-load sorted leaves by detecting ordered input, retaining the first key and last value, and merging batches; use plain `conj` for eight or fewer entries.
3. **Study memory behavior.** Determine which layouts suit Go's memory model: per-node allocation versus contiguous leaves, pointer-heavy tries versus chunked arrays, and arena or bump allocation for builders. Analyze escape behavior and profile with pprof or massif. For sorted maps and sets, use a persistent B+tree (16-entry leaves, 32-child inner nodes) rather than a red-black tree; see <https://yogthos.net/posts/2026-11-23-b+trees.html>. Always measure small, medium, and large inputs, since memory layout dominates performance.
4. **Explore goroutines and channels.** Investigate how the collection API fits Go's concurrency model.

## Non-goals

- Runtime laziness, thunks, and coinductive streams.
- Zero *total* allocation: result structures are still allocated.

I may revisit laziness, since the transducer model can support both push- and pull-based evaluation.

## License

MIT, the same license as [Lisette](https://lisette.run). See [LICENSE](LICENSE).
