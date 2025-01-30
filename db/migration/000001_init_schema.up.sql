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
  "macaddress" varchar NOT NULL,
  "ipaddress" varchar NOT NULL,
  "mtu" bigint,
  "name" varchar,
  "type" interface_types NOT NULL,
  "tag" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE TABLE "vxlan_tunnel" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "tag" bigint UNIQUE NOT NULL,
  "tunnel_ip" varchar UNIQUE NOT NULL,
  "local_ip" varchar NOT NULL,
  "remote_ip" varchar NOT NULL,
  "remote_mac" varchar NOT NULL,
  "phys_iface" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bgp_peer" (
  "id" bigint PRIMARY KEY,
  "name" varchar NOT NULL,
  "as_no" int NOT NULL,
  "neighbor_address" int NOT NULL,
  "local_as" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bgp_advertisement" (
  "id" bigint PRIMARY KEY,
  "name" varchar NOT NULL,
  "destination_cidr" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "system" (
   "id"  bigserial PRIMARY KEY,
   "component" varchar NOT NULL,
   "initialised" BOOLEAN NOT NULL DEFAULT  FALSE
);

ALTER TABLE "vxlan_tunnel" ADD FOREIGN KEY ("phys_iface") REFERENCES "interfaces" ("id");
