create table advertisement
(
    id          serial not null constraint advertisement_pkey primary key,
    name        varchar(200),
    description varchar(1000),
    images_urls text[3],
    price       integer,
    date        timestamp
);

INSERT INTO Advertisement
  (name, description,images_urls, price,date)
VALUES
  ('flower','red flower','{"link1", "link2"}',5,now()), 
  ('flower2','green flower','{"link21"}',4,now() - INTERVAL '10 DAYS'), 
  ('dog','big dog','{"link","linktodog","linklink"}',55,now() - INTERVAL '150 DAYS'),
  ('table','round table','{}',45,now() - INTERVAL '13 DAYS'),
  ('table 2','new table','{"link2323","url"}',64,now() - INTERVAL '1 DAYS'),
  ('pen 2','new pen','{"link23"}',4,now() - INTERVAL '6 DAYS'),
  ('red pen','red pen','{}',1,now() - INTERVAL '588 DAYS'),
  ('pot','old pot','{"url1", "url2","url3"}',20,now() - INTERVAL '88 DAYS'),
  ('carpet','old carpet','{"url331", "url222","ur2l3"}',222,now() - INTERVAL '8 DAYS'),
  ('lamp','old lamp','{"url331e"}',22,now() - INTERVAL '28 DAYS'),
  ('cat','just cat','{"url3"}',2,now() - INTERVAL '48 DAYS'),
  ('dog','good boy','{"url3ee","url666"}',0,now() - INTERVAL '4 DAYS'),
  ('door','good door','{"url3ewe","url66w6","linklinklink"}',56,now() - INTERVAL '14 DAYS'),
  ('teapot','good teapot','{"urle","urlw6","linklinklink2"}',26,now() - INTERVAL '37 DAYS'),
  ('refrigerator','white refrigerator','{}',350,now() - INTERVAL '3 DAYS'),
  ('bed','big bed','{}',250,now() - INTERVAL '23 DAYS'),
  ('knife','sharp knife','{"url2222"}',70,now() - INTERVAL '66 DAYS'),
  ('cupboard','big cupboard','{}',98,now() - INTERVAL '40 DAYS'),
  ('wardrobe','roomy wardrobe','{"link11","link32"}',157,now() - INTERVAL '140 DAYS'),
  ('piano','old piano','{"linkk","llink","linnk"}',207,now() - INTERVAL '10 DAYS'),
  ('garage','best garage','{"url222"}',1007,now() - INTERVAL '64 DAYS'),
  ('house','house in town','{"linktophoto","link2tophoto"}',10007,now() - INTERVAL '128 DAYS'),
  ('TV','cheap TV','{"linktophoto2"}',77,now() - INTERVAL '18 DAYS'),
  ('TV','expensive TV','{"linktophoto22222"}',250,now() - INTERVAL '77 DAYS');





