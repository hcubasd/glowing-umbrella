import pg from "pg";

// sales.crm_deals.amount and sales.crm_products.price are `numeric` columns (OID 1700).
// node-postgres returns those as strings by default to avoid float precision loss;
// the API contract (matching the old Go/lib-pq behavior) is a JSON number.
pg.types.setTypeParser(1700, (value) => (value === null ? null : Number.parseFloat(value)));

// sales.crm_deals.expected_close_date is a `date` column (OID 1082). node-postgres's default
// parser builds a local-timezone Date from date-only values, which can shift the calendar day
// depending on the container's TZ. Keep the raw "YYYY-MM-DD" string instead.
pg.types.setTypeParser(1082, (value) => value);

export const pool = new pg.Pool();
