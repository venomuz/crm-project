CREATE TABLE "companies"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" TEXT NOT NULL,
    "info" TEXT,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);
CREATE TABLE "payment"(
    "id" UUID NOT NULL PRIMARY KEY,
    "company_id" UUID NOT NULL,
    "paid_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "expired_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "payment_method" TEXT NOT NULL,
    FOREIGN KEY("company_id") REFERENCES "companies"("id") ON DELETE CASCADE
);
CREATE TABLE "users"(
    "id" UUID NOT NULL PRIMARY KEY,
    "company_id" UUID NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" TEXT NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "role" VARCHAR(255) NOT NULL,
    "gender" VARCHAR(255) NOT NULL,
    "info" TEXT,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE,
    "joined_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "left_date" TIMESTAMP(0) WITHOUT TIME ZONE,
    FOREIGN KEY("company_id") REFERENCES "companies"("id") ON DELETE CASCADE
);
CREATE TABLE "groups"(
    "id" UUID NOT NULL PRIMARY KEY,
    "company_id" UUID ,
    "user_id" UUID ,
    "group_name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY("company_id") REFERENCES "companies"("id") ON DELETE CASCADE,
    FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);  

CREATE TABLE "schedule_student"(
    "id" UUID NOT NULL PRIMARY KEY,
    "student_id" UUID NOT NULL,
    "group_id" UUID NOT NULL,
    "attendance" BOOLEAN NOT NULL,
    "attendance_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "feedback" TEXT,
    FOREIGN KEY("group_id") REFERENCES "groups"("id") ON DELETE CASCADE
);

CREATE TABLE "lesson"(
    "id" UUID NOT NULL PRIMARY KEY,
    "group_id" UUID NOT NULL,
    "date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    FOREIGN KEY("group_id") REFERENCES "groups"("id") ON DELETE CASCADE
);