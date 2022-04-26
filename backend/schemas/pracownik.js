import { Sequelize, Model, DataTypes } from 'sequelize';

const Pracownik = sequelize.define('Pracownik', {
  id: {type: DataTypes.INTEGER, primaryKey: true},
  PESEL: DataTypes.STRING,
  imie: DataTypes.STRING,
  nazwisko: DataTypes.STRING,
  login: DataTypes.STRING,
  haslo: DataTypes.STRING,
});

module.exports = Pracownik