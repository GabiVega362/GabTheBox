CREATE OR REPLACE VIEW "images" AS
    SELECT
        "id" AS "lab",
        "org" || '/' || "environment" || ':' || "release" AS "image"
    FROM "labs";

CREATE OR REPLACE VIEW "user_labs" AS
SELECT
    L."id" AS "lab",
    L."description",
    L."org" || '/' || L."environment" || ':' || L."release" AS "image",
    UL."user",
    "port"
FROM "labs" L LEFT JOIN "users_labs" UL ON L."id" = UL."lab";