

SELECT
    s.store_id,
    SUM(s.quantity),
    JSONB_AGG (
        JSONB_BUILD_OBJECT (
            'product_id', p.product_id,
            'product_name', p.product_name,
            'brand_id', p.brand_id,
            'category_id', p.category_id,
            'model_year', p.model_year,
            'list_price', p.list_price,
            'quantity', s.quantity
        )
    ) AS product_data
FROM stocks AS s
LEFT JOIN products AS p ON p.product_id = s.product_id
WHERE s.store_id = 1
GROUP BY s.store_id


CREATE INDEX stock_product_idx ON stocks (product_id);


WITH order_item_data AS (
    SELECT
        oi.order_id AS order_id,
        JSONB_AGG (
            JSONB_BUILD_OBJECT (
                'order_id', oi.order_id,
                'item_id', oi.item_id,
                'product_id', oi.product_id,
                'quantity', oi.quantity,
                'list_price', oi.list_price,
                'discount', oi.discount
            )
        ) AS order_items

    FROM order_items AS oi
    WHERE oi.order_id = 1616
    GROUP BY oi.order_id
)
SELECT
    o.order_id, 
    o.customer_id,

    c.customer_id,
    c.first_name,
    c.last_name,
    COALESCE(c.phone, ''),
    c.email,
    COALESCE(c.street, ''),
    COALESCE(c.city, ''),
    COALESCE(c.state, ''),
    COALESCE(c.zip_code, 0),
    
    o.order_status,
    CAST(o.order_date::timestamp AS VARCHAR),
    CAST(o.required_date::timestamp AS VARCHAR),
    COALESCE(CAST(o.shipped_date::timestamp AS VARCHAR), ''),
    o.store_id,

    s.store_id,
    s.store_name,
    COALESCE(s.phone, ''),
    COALESCE(s.email, ''),
    COALESCE(s.street, ''),
    COALESCE(s.city, ''),
    COALESCE(s.state, ''),
    COALESCE(s.zip_code, ''),

    o.staff_id,
    st.staff_id,
    st.first_name,
    st.last_name,
    st.email,
    COALESCE(st.phone, ''),
    st.active,
    st.store_id,
    COALESCE(st.manager_id, 0),

    oi.order_items

FROM orders AS o
JOIN customers AS c ON c.customer_id = o.customer_id
JOIN stores AS s ON s.store_id = o.store_id
JOIN staffs AS st ON st.staff_id = o.staff_id
JOIN order_item_data AS oi ON oi.order_id = o.order_id
WHERE o.order_id = 1616



SELECT 
    COUNT(*) OVER(),
    s.first_name || ' ' || s.last_name AS staff_name,
    c.category_name,
    p.product_name,
    ori.quantity,
    ori.list_price,
    o.order_date
FROM
    orders AS o
JOIN order_items AS ori ON o.order_id = ori.order_id
JOIN products AS p ON p.product_id = ori.product_id
JOIN categories AS c ON c.category_id = p.category_id
JOIN staffs AS s ON s.staff_id = o.staff_id
WHERE o.order_status = 4
;


SELECT
    o.order_id,
    oi.quantity,
    oi.list_price
FROM
    orders AS o
JOIN order_items AS oi ON o.order_id = oi.order_id
WHERE o.order_status = 4 AND o.order_id = 1




SELECT
    p.product_name

FROM
    stocks AS st
JOIN stores AS s ON s.store_id = st.store_id
JOIN products AS p ON p.product_id = st.product_id