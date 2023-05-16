

  WITH emax AS(SELECT event_type, max(time) as maxt
    FROM events
    GROUP BY event_type
    HAVING count(event_type) > 1
    ORDER BY event_type ASC
  )
  SELECT result.event_type, result.frow - result.srow AS value
  FROM (
    SELECT emax.event_type,
    (SELECT events.value FROM events WHERE events.event_type = emax.event_type AND events.time = emax.maxt) as frow,
    (SELECT events.value FROM events WHERE events.event_type = emax.event_type AND events.time <> emax.maxt ORDER BY time DESC LIMIT 1) as srow
  FROM emax
  ) as result
