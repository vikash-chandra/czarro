// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs
// vikash1111@mailinator.com 

Table cz_country {
  id serial [primary key]
  iso varchar(2) [not null]
  name varchar(80) [not null]
  nicename varchar(80) [not null]
  iso3 varchar(3) [not null]
  numcode integer [not null,default:0]
  phone_code integer [not null,default:0]
}


Table cz_roles {
  id serial [primary key]
  role_name varchar(20) [not null]
  status_id integer [not null]
  create_user bigserial [not null]
  modify_user bigserial [not null]
  created_at timestamptz [not null, default: `now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']
  visible boolean [not null, default:true]
}

Table cz_users {
  id bigserial [not null]
  unique_id uuid [primary key]
  role_id integer [not null,ref: > cz_roles.id]
  first_name varchar [not null]
  middle_name varchar [not null]
  last_name varchar [not null]
  dob timestamptz [not null, default:'0001-01-01']
  country_code integer [not null,ref: > cz_country.id]
  phone varchar [not null]
  email varchar [not null]
  salt varchar [not null]
  password varchar [not null]
  status_id integer [not null]
  create_user bigserial [not null]
  modify_user bigserial [not null]
  created_at timestamptz [not null, default:`now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']

  Indexes {
    unique_id
    phone
  }
}

Table cz_users_address {
  id bigserial [primary key]
  user_id uuid [not null, ref: > cz_users.unique_id]
  country_code integer [not null,ref: > cz_country.id]
  address1 varchar(100) [not null]
  address2 varchar(100) [not null]
  address3 varchar(100) [not null]
  address4 varchar(100) [not null]
  location varchar(100) [not null]
  status_id integer [not null]
  create_user bigserial [not null]
  modify_user bigserial [not null]
  created_at timestamptz [not null, default:`now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']

  Indexes {
    user_id
  }
}


Table cz_notification {
  id bigserial [primary key]
  status_id integer [not null,default:0]
  create_user bigserial [not null]
  modify_user bigserial [not null]
  created_at timestamptz [not null, default: `now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']
}


Table cz_services {
  id serial [primary key]
  title varchar(50) [not null]
  short_name varchar(30) [not null]
  description text [not null]
  send_notification integer [not null]
  status_id integer [not null]
  create_user bigserial [not null]
  modify_user bigserial [not null]
  created_at timestamptz [not null, default: `now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']
}


