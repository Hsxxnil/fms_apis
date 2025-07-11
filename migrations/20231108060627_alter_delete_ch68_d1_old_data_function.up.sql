drop function delete_ch68_d1_old_data();

CREATE OR REPLACE FUNCTION delete_poison_old_data()
    RETURNS SETOF poison_data as
$$
DECLARE
    one_week_ago_date DATE;
    current_sid       text;
BEGIN
    FOR current_sid IN SELECT DISTINCT sid FROM poison_data
        LOOP
            SELECT MAX(date_time) - INTERVAL '7 days'
            INTO one_week_ago_date
            FROM poison_data
            WHERE sid = current_sid;

            -- Delete old data
            DELETE
            FROM poison_data
            WHERE sid = current_sid
              AND date_time <= one_week_ago_date;
        END LOOP;
END;
$$
    LANGUAGE plpgsql;