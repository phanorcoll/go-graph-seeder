g.V().hasLabel("address").
  group().
    by(
      project('a1', 'a2', 'city', 'state', 'zip', 'country').
        by(coalesce(values('address_1'), constant(''))).
        by(coalesce(values('address_2'), constant(''))).
        by(coalesce(values('city'), constant(''))).
        by(coalesce(values('state'), constant(''))).
        by(coalesce(values('postal_code'), constant(''))).
        by(coalesce(values('country'), constant('')))
    ).
    by(
      order(local).by('created_at', asc).fold()
    ).
  unfold().
  filter(select(values).count(local).is(gt(1))).
  select(values).
  range(local, 1, -1).
  unfold().
  sideEffect(inE().drop()).
  drop()
