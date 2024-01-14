// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs
// vikash1111@mailinator.com 

Table cz_status {
  id serial [primary key]
  status varchar(50) [not null]
  description varchar(100) [not null, default:'']
  visible boolean [not null, default:false]
}

Table cz_currency {
  id serial [primary key]
  currency varchar(50) [not null, default:'INR']
  description varchar(100) [not null, default:'']
  visible boolean [not null, default:false]
}

Table cz_country {
  id serial [primary key]
  iso varchar(2) [not null]
  name varchar(80) [not null]
  nicename varchar(80) [not null]
  iso3 varchar(3) [not null]
  numcode varchar(6) [not null,default:'']
  phone_code integer [not null,default:0]
}


Table cz_roles {
  id serial [primary key]
  role_name varchar(20) [not null]
  status_id integer [not null]
  create_user bigint [not null,default:0]
  modify_user bigint [not null,default:0]
  created_at timestamptz [not null, default: `now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']
  visible boolean [not null, default:true]
}

Table cz_users {
  id bigserial [primary key]
  unique_id uuid [not null]
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
  status_id integer [not null, ref: > cz_status.id]
  create_user bigint [not null,default:0]
  modify_user bigint [not null,default:0]
  created_at timestamptz [not null, default:`now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']

  Indexes {
    unique_id
    phone
  }
}

Table cz_users_address {
  id bigserial [primary key]
  user_id bigserial [not null, ref: > cz_users.id]
  country_code integer [not null,ref: > cz_country.id]
  address1 varchar(100) [not null]
  address2 varchar(100) [not null]
  address3 varchar(100) [not null]
  address4 varchar(100) [not null]
  location varchar(100) [not null]
  status_id integer [not null]
  create_user bigint [not null,default:0]
  modify_user bigint [not null,default:0]
  created_at timestamptz [not null, default:`now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']

  Indexes {
    user_id
  }
}

Table cz_products {
  id serial [primary key]
  title varchar(50) [not null]
  short_name varchar(30) [not null]
  description text [not null]
  sms_noti boolean [not null, default:false]
  email_noti boolean [not null, default:false]
  call_noti boolean [not null, default:false]
  image varchar(100) [not null, default:'']
  currency_id integer [not null, ref: > cz_currency.id]
  price float [not null, default:0]
  status_id integer [not null, default:0]
  create_user bigint [not null,default:0]
  modify_user bigint [not null,default:0]
  created_at timestamptz [not null, default: `now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']

   Indexes {
    id
  }
}
