g.V().hasLabel("address").
  group().
    by(
      // Define the grouping key based on address properties
      project('a1', 'a2', 'city', 'state', 'zip', 'country').
        by(coalesce(values('address_1'), constant(''))).
        by(coalesce(values('address_2'), constant(''))).
        by(coalesce(values('city'), constant(''))).
        by(coalesce(values('state'), constant(''))).
        by(coalesce(values('postal_code'), constant(''))).
        by(coalesce(values('country'), constant('')))
    ).
    by(
      // For each group key, get the list of vertices
      order(local).by('created_at', asc).fold()
    ).
  unfold(). // Unfold the map entries (addressKey -> List<Vertex>)
  filter(select(values).count(local).is(gt(1))). // Keep only entries where the list has more than one vertex
  select(values). // Get the list of duplicate vertices (survivor + victims), ordered by created_at
  range(local, 1, -1). // *** MODIFIED: Get the victims from each list ***
  unfold(). // *** MODIFIED: Unfold the list of victims into a stream of individual vertices ***
  count() // *** MODIFIED: Count the total number of victim vertices in the stream ***
