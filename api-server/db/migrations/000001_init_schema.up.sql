CREATE TYPE "interface_types" AS ENUM (
  'vlan',
  'vxlan',
  'bond',
  'ether',
  'unknown',
  'loopback'
);

CREATE TABLE "interfaces" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "macaddress" varchar NOT NULL,
  "ipaddress" varchar NOT NULL,
  "mtu" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "vlans" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "ipaddress" varchar NOT NULL,
  "netmask" varchar NOT NULL,
  "lower" bigserial NOT NULL,
  "tag" bigint NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "vxlan_tunnel" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "tag" bigint UNIQUE NOT NULL,
  "tunnel_ip" varchar UNIQUE NOT NULL,
  "local_ip" varchar NOT NULL,
  "remote_ip" varchar NOT NULL,
  "remote_mac" varchar NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE TABLE "bgp_peer" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "as_no" bigint NOT NULL UNIQUE,
  "neighbor_address" varchar NOT NULL,
  "local_as" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bgp_advertisement" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "destination_cidr" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "system" (
   "id"  bigserial PRIMARY KEY,
   "component" varchar NOT NULL,
   "initialised" BOOLEAN NOT NULL DEFAULT  FALSE
);

ALTER TABLE "vlans" ADD FOREIGN KEY ("lower") REFERENCES "interfaces" ("id");
-- ALTER TABLE "vxlan_tunnel" ADD FOREIGN KEY ("phys_iface") REFERENCES "interfaces" ("id");
