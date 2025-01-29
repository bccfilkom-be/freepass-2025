"use strict";
const bcrypt = require("bcrypt");

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, _Sequelize) {
    const hashedPassword1 = await bcrypt.hash("Str0ngP@assw0rd1", 10);
    const hashedPassword2 = await bcrypt.hash("Str0ngP@assw0rd2", 10);
    const hashedPassword3 = await bcrypt.hash("Str0ngP@assw0rd3", 10);

    return queryInterface.bulkInsert(
      "Users",
      [
        {
          id: "d597d69f-bea7-4951-9ba9-0bc330fd16cd",
          username: "Amber",
          email: "amber@example.com",
          password: hashedPassword1,
          role: "user",
          createdAt: new Date(),
          updatedAt: new Date(),
        },
        {
          id: "8013964b-07f9-43b8-87cc-4b6b748313af",
          username: "Lucia",
          email: "lucia@example.com",
          password: hashedPassword2,
          role: "user",
          createdAt: new Date(),
          updatedAt: new Date(),
        },
        {
          id: "b8d6c1e4-7e6b-4b7f-9f8c-5e6c7b2e1b4b",
          username: "Bianca",
          email: "bianca@example.com",
          password: hashedPassword3,
          role: "coordinator",
          createdAt: new Date(),
          updatedAt: new Date(),
        },
      ],
      {}
    );
  },

  async down(queryInterface, _Sequelize) {
    return queryInterface.bulkDelete("Users", null, {});
  },
};
