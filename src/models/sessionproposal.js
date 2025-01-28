"use strict";
const { Model } = require("sequelize");
module.exports = (sequelize, DataTypes) => {
  class SessionProposal extends Model {
    static associate(models) {
      this.belongsTo(models.User, { foreignKey: "user_id" });
    }
  }
  SessionProposal.init(
    {
      id: {
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
        type: DataTypes.INTEGER,
      },
      user_id: {
        type: DataTypes.UUID,
        allowNull: false,
        references: {
          model: "User",
          key: "id",
        },
      },
      title: {
        type: DataTypes.STRING,
        allowNull: false,
      },
      description: {
        type: DataTypes.TEXT,
        allowNull: false,
      },
      start_time: {
        type: DataTypes.DATE,
        allowNull: false,
      },
      end_time: {
        type: DataTypes.DATE,
        allowNull: false,
      },
      status: {
        type: DataTypes.ENUM("pending", "accepted", "rejected"),
        allowNull: false,
        defaultValue: "pending",
      },
      available_seats: {
        type: DataTypes.INTEGER,
        allowNull: false,
        defaultValue: 20,
      },
    },
    {
      sequelize,
      modelName: "SessionProposal",
    }
  );
  return SessionProposal;
};
