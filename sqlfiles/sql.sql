-- Generated by Oracle SQL Developer Data Modeler 4.1.5.907
--   at:        2022-04-08 21:04:46 CEST
--   site:      Oracle Database 12c
--   type:      Oracle Database 12c

-- test insert
-- INSERT INTO administracja (pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id) VALUES ('00000000000', 'a', 'b', 'c', 'd', 1);
-- INSERT INTO pracownicy (pesel, imie, nazwisko,) VALUES ('00000000000', 'a', 'b');

-- check
CREATE TABLE Kategoria_Prawa_Jazdy
  (
    kategoria   VARCHAR (4000) ,
    id 		SERIAL PRIMARY KEY
  ) ;

-- check
CREATE TABLE Kurs
  (
    data_rozpoczecia TIMESTAMP ,
    data_zakonczenia TIMESTAMP ,
    czas_przejazdu FLOAT ,
    ladunek FLOAT ,
    Trasa_ID       INTEGER NOT NULL ,
    Kierowcy_ID INTEGER NOT NULL ,
    Pojazd_ID     INTEGER NOT NULL
  ) ;
ALTER TABLE Kurs ADD CONSTRAINT Kurs_PK PRIMARY KEY ( Trasa_ID, Kierowcy_ID, Pojazd_ID ) ;

-- check
CREATE TABLE Marka
  ( 
    nazwa 	VARCHAR (4000) ,
    id 		SERIAL PRIMARY KEY
  ) ;

-- check
CREATE TABLE Pojazd
  (
    numer_silnika       INTEGER ,
    pojemnosc_silnika   INTEGER ,
    Marka_ID           INTEGER NOT NULL ,
    id    		SERIAL PRIMARY KEY
  ) ;

-- check
CREATE TABLE Pracownicy
  (
    id       SERIAL PRIMARY KEY ,
    PESEL    VARCHAR (11) ,
    imie     VARCHAR (4000) ,
    nazwisko VARCHAR (4000)
  ) ;

-- check
CREATE TABLE Kategoria_Kierowcy
  (
    Kierowcy_ID INTEGER NOT NULL ,
    Kategoria_Prawa_Jazdy_ID INTEGER NOT NULL
  ) ;
ALTER TABLE Kategoria_Kierowcy ADD CONSTRAINT Kategoria_Kierowcy_PK PRIMARY KEY ( Kierowcy_ID, Kategoria_Prawa_Jazdy_ID ) ;

-- check
CREATE TABLE Stanowisko_Administracyjne
  (
    typ         VARCHAR (4000) ,
    id 		SERIAL PRIMARY KEY
  ) ;

-- check
CREATE TABLE Trasa
  (
    miejscowosc_poczatkowa 	VARCHAR (4000) ,
    miejscowosc_koncowa    	VARCHAR (4000) ,
    id              		SERIAL PRIMARY KEY
  ) ;

-- check
CREATE TABLE Administracja
  (
    Administracja_ID SERIAL PRIMARY KEY ,
    login         VARCHAR (4000) ,
    haslo         VARCHAR (4000) ,
    Stanowisko_Administracyjne_ID INTEGER NOT NULL
  ) INHERITS (Pracownicy) ;

-- check
CREATE TABLE Kierowcy
  (
    Kierowcy_ID SERIAL PRIMARY KEY
  ) INHERITS (Pracownicy) ;

-- check
CREATE TABLE Pojazd_Ciezarowy
  (
    ladownosc FLOAT
  ) INHERITS(Pojazd) ;


ALTER TABLE Administracja ADD CONSTRAINT Administracja_Stanowisko_Administracyjne_FK FOREIGN KEY ( Stanowisko_Administracyjne_ID ) REFERENCES Stanowisko_Administracyjne ( ID ) ;

ALTER TABLE Kategoria_Kierowcy ADD CONSTRAINT FK_ASS_2 FOREIGN KEY ( Kierowcy_ID ) REFERENCES Kierowcy ( Kierowcy_ID ) ;

ALTER TABLE Kategoria_Kierowcy ADD CONSTRAINT FK_ASS_3 FOREIGN KEY ( Kategoria_Prawa_Jazdy_ID ) REFERENCES Kategoria_Prawa_Jazdy ( ID ) ;

ALTER TABLE Kurs ADD CONSTRAINT Kurs_Kierowcy_FK FOREIGN KEY ( Kierowcy_ID ) REFERENCES Kierowcy ( Kierowcy_ID ) ;

ALTER TABLE Kurs ADD CONSTRAINT Kurs_Pojazd_FK FOREIGN KEY ( Pojazd_ID ) REFERENCES Pojazd ( ID ) ;

ALTER TABLE Kurs ADD CONSTRAINT Kurs_Trasa_FK FOREIGN KEY ( Trasa_ID ) REFERENCES Trasa ( ID ) ;

ALTER TABLE Pojazd_Ciezarowy ADD CONSTRAINT Pojazd_Ciezarowy_Pojazd_FK FOREIGN KEY ( ID ) REFERENCES Pojazd ( ID ) ;

ALTER TABLE Pojazd ADD CONSTRAINT Pojazd_Marka_FK FOREIGN KEY ( Marka_ID ) REFERENCES Marka ( ID ) ;



--CREATE SEQUENCE Administracja_ID_SEQ
--   OWNED BY Administracja.Administracja_ID;
--
--ALTER TABLE Administracja
--   ALTER Administracja_ID
--      SET DEFAULT nextval('Administracja_ID_SEQ'::regclass);

--CREATE SEQUENCE Kierowcy_ID_SEQ
--   OWNED BY Kierowcy.Kierowcy_ID;  

--ALTER TABLE Kierowcy
--   ALTER Kierowcy_ID
--      SET DEFAULT nextval('Kierowcy_ID_SEQ'::regclass);

--CREATE SEQUENCE Kategoria_Prawa_Jazdy_ID_SEQ
--   OWNED BY Kategoria_Prawa_Jazdy.id;

--ALTER TABLE Kategoria_Prawa_Jazdy
--   ALTER id
--      SET DEFAULT nextval('Kategoria_Prawa_Jazdy_ID_SEQ'::regclass);


--CREATE SEQUENCE Kierowcy_ID_SEQ
--   OWNED BY Kierowcy.id;

--ALTER TABLE Kierowcy
--   ALTER id
--      SET DEFAULT nextval('Kierowcy_ID_SEQ'::regclass);


--CREATE SEQUENCE Marka_ID_SEQ
--   OWNED BY Kierowcy.id;

--ALTER TABLE Marka
--   ALTER id
--      SET DEFAULT nextval('Marka_ID_SEQ'::regclass);



--CREATE SEQUENCE Pojazd_ID_SEQ
--   OWNED BY Pojazd.id;

--ALTER TABLE Pojazd
--   ALTER id
--      SET DEFAULT nextval('Pojazd_ID_SEQ'::regclass);


--CREATE SEQUENCE Stanowisko_Administracyjne_ID_SEQ
--   OWNED BY Stanowisko_Administracyjne.id;

--ALTER TABLE Stanowisko_Administracyjne
--   ALTER id
--      SET DEFAULT nextval('Stanowisko_Administracyjne_ID_SEQ'::regclass);


--CREATE SEQUENCE Trasa_ID_SEQ
--   OWNED BY Trasa.id;

--ALTER TABLE Trasa
--   ALTER id
--      SET DEFAULT nextval('Trasa_ID_SEQ'::regclass);


-- Oracle SQL Developer Data Modeler Summary Report: 
-- 
-- CREATE TABLE                            11
-- CREATE INDEX                             3
-- ALTER TABLE                             18
-- CREATE VIEW                              0
-- ALTER VIEW                               0
-- CREATE PACKAGE                           0
-- CREATE PACKAGE BODY                      0
-- CREATE PROCEDURE                         0
-- CREATE FUNCTION                          0
-- CREATE TRIGGER                           6
-- ALTER TRIGGER                            0
-- CREATE COLLECTION TYPE                   0
-- CREATE STRUCTURED TYPE                   0
-- CREATE STRUCTURED TYPE BODY              0
-- CREATE CLUSTER                           0
-- CREATE CONTEXT                           0
-- CREATE DATABASE                          0
-- CREATE DIMENSION                         0
-- CREATE DIRECTORY                         0
-- CREATE DISK GROUP                        0
-- CREATE ROLE                              0
-- CREATE ROLLBACK SEGMENT                  0
-- CREATE SEQUENCE                          6
-- CREATE MATERIALIZED VIEW                 0
-- CREATE SYNONYM                           0
-- CREATE TABLESPACE                        0
-- CREATE USER                              0
-- 
-- DROP TABLESPACE                          0
-- DROP DATABASE                            0
-- 
-- REDACTION POLICY                         0
-- TSDP POLICY                              0
-- 
-- ORDS DROP SCHEMA                         0
-- ORDS ENABLE SCHEMA                       0
-- ORDS ENABLE OBJECT                       0
-- 
-- ERRORS                                   3
-- WARNINGS                                 0