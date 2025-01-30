CREATE TYPE "interface_types" AS ENUM (
  'vlan',
  'vxlan',
  'bond'
);
CREATE TABLE "system" {
   "id"  bigserial PRIMARY KEY
   "component" varchar2 NOT NULL
   "initialised" BOOLEAN DEFAULT FALSE
}

CREATE TABLE "interfaces" (
  "id" bigserial PRIMARY KEY,
  "macaddress" varchar2 NOT NULL,
  "ipaddress" varchar2,
  "mtu" int,
  "name" varchar2,
  "type" interface_types NOT NULL,
  "vlan_profile" bigint,
  "vxlan_profile" bigint,
  "bond_profile" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bond_profile" (
  "id" bigint PRIMARY KEY
);

CREATE TABLE "vlan_profile" (
  "id" bigserial PRIMARY KEY,
  "name" varchar2,
  "tag" bigint UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "vxlan_tunnel" (
  "id" bigserial PRIMARY KEY,
  "name" varchar2,
  "tag" bigint UNIQUE NOT NULL,
  "tunnel_ip" varchar2 UNIQUE NOT NULL,
  "local_ip" varchar2 NOT NULL,
  "remote_ip" varchar2 NOT NULL,
  "remote_mac" varchar2 NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bgp_peer" (
  "id" bigint PRIMARY KEY,
  "as" int NOT NULL,
  "neighbor_address" int NOT NULL,
  "local_as" int,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bgp_advertisement" (
  "id" bigint PRIMARY KEY,
  "destination" varchar2 NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "interfaces" ADD FOREIGN KEY ("vlan_profile") REFERENCES "vlan_profile" ("id");

ALTER TABLE "interfaces" ADD FOREIGN KEY ("vxlan_profile") REFERENCES "vxlan_tunnel" ("id");

ALTER TABLE "interfaces" ADD FOREIGN KEY ("bond_profile") REFERENCES "bond_profile" ("id");
