SELECT eventType
FROM events
GROUP BY eventType
HAVING count(*)>1000;

SELECT *
FROM events
WHERE toStartOfMonth(eventTime) = eventTime;

SELECT userID
FROM (
    SELECT userID, count(DISTINCT eventType) as eventTypeCount
    FROM events
    GROUP BY userID
) AS UserEventCounts
WHERE eventTypeCount > 3;


