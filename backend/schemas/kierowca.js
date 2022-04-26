import { Sequelize, Model, DataTypes } from 'sequelize';

const Kierowca = sequelize.define('Kierowcy', {
  Kierowcy_ID: {type: DataTypes.INTEGER, primaryKey: true},
  id: DataTypes.INTEGER,
  PESEL: DataTypes.STRING,
  imie: DataTypes.STRING,
  nazwisko: DataTypes.STRING,
  login: DataTypes.STRING,
  haslo: DataTypes.STRING,
});

module.exports = Kierowca