-- Create a keyspace
CREATE KEYSPACE IF NOT EXISTS eos WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

-- Create a table
CREATE TABLE IF NOT EXISTS eos.lifx (
    name varchar,
    selector varchar,
    ip varchar,
    zone int,
    scene varchar,
    powered varchar,
    hue float,
    saturation float,
    kelvin int,
    brightness float,
    PRIMARY KEY ((scene), selector, zone)
);

-- Create a table
CREATE TABLE IF NOT EXISTS eos.hue (
    name varchar,
    selector int,
    scene varchar,
    powered boolean,
    color_x float,
    color_y float,
    brightness int,
    PRIMARY KEY ((scene), selector)
);