"use strict";
const bcrypt = require("bcrypt");

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, _Sequelize) {
    const hasedPassword1 = await bcrypt.hash("Str0ngP@assw0rd1", 10);
    const hasedPassword2 = await bcrypt.hash("Str0ngP@assw0rd2", 10);

    return queryInterface.bulkInsert(
      "Users",
      [
        {
          id: "d597d69f-bea7-4951-9ba9-0bc330fd16cd",
          username: "Amber",
          email: "amber@example.com",
          password: hasedPassword1,
          createdAt: new Date(),
          updatedAt: new Date(),
        },
        {
          id: "8013964b-07f9-43b8-87cc-4b6b748313af",
          username: "Lucia",
          email: "lucia@example.com",
          password: hasedPassword2,
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
