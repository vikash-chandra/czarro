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

Table cz_vendors {
  id bigserial [primary key]
  vendor_id uuid [not null]
  vendor_name varchar(150) [not null]
  registration_number varchar(50) [not null]
  website_url varchar(100) 
  contact_email varchar(100) [not null]
  contact_number varchar(20) [not null]
  salt varchar [not null]
  password varchar [not null]
  country_code integer [not null,ref: > cz_country.id]
  status_id integer [not null, ref: > cz_status.id]
  create_user bigint [not null,default:0]
  modify_user bigint [not null,default:0]
  created_at timestamptz [not null, default:`now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']

  Indexes {
    vendor_id
    vendor_name
  }
}

Table cz_vendors_shops {
  id bigserial [primary key]
  vendor_id bigserial [not null, ref: > cz_vendors.id]
  shop_name varchar(200) [not null]
  address varchar(250) [not null]
  city varchar(100) [not null]
  state varchar(100) [not null]
  postal_code varchar(20) [not null]
  country_code integer [not null,ref: > cz_country.id]
  location varchar(100) [not null]
  create_user bigint [not null,default:0]
  modify_user bigint [not null,default:0]
  created_at timestamptz [not null, default:`now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']

  Indexes {
    shop_name
    city
  }
}

Table cz_vendors_workers {
  id bigserial [primary key]
  worker_id uuid [not null]
  shop_id bigserial [not null, ref: > cz_vendors_shops.id]
  first_name varchar(50) [not null]
  middle_name varchar(50) 
  last_name varchar(50) [not null]
  designation varchar(50) [not null]
  contact_number varchar(20) [not null]
  contact_email varchar(100)
  adhaar_card_number varchar(20) [not null]
  dob timestamptz [not null, default:'0001-01-01']
  create_user bigint [not null,default:0]
  modify_user bigint [not null,default:0]
  created_at timestamptz [not null, default:`now()`]
  modified_at timestamptz [not null, default:'0001-01-01 00:00:00+00']
  //CONSTRAINT fk_shop_id FOREIGN KEY (shop_id) REFERENCES shops(shop_id)

  Indexes {
    worker_id
    contact_number
  }
}