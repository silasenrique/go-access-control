CREATE TABLE COMPANY (
  Id             serial primary key,
	Code           varchar(10) unique not null,
	Name           varchar(100),
	SiteUrl        varchar(200),
	CreationDate   date,
	LastChangeDate date
);