drop function delete_poison_old_data();

CREATE OR REPLACE FUNCTION delete_ch68_d1_old_data()
    RETURNS SETOF ch68_d1_data as
$$
DECLARE
    one_week_ago_date DATE;
    current_sid       text;
BEGIN
    FOR current_sid IN SELECT DISTINCT sid FROM ch68_d1_data
        LOOP
            SELECT MAX(date_time) - INTERVAL '7 days'
            INTO one_week_ago_date
            FROM ch68_d1_data
            WHERE sid = current_sid;

            -- Delete old data
            DELETE
            FROM ch68_d1_data
            WHERE sid = current_sid
              AND date_time <= one_week_ago_date;
        END LOOP;
END;
$$
    LANGUAGE plpgsql;