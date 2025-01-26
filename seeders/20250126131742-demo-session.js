"use strict";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, _Sequelize) {
    return queryInterface.bulkInsert(
      "Sessions",
      [
        {
          title: "Session 1",
          description: "Description for session 1",
          start_time: new Date(),
          end_time: new Date(new Date().getTime() + 60 * 60 * 1000), // 1 hour later
          available_seats: 10,
          createdAt: new Date(),
          updatedAt: new Date(),
        },
        {
          title: "Session 2",
          description: "Description for session 2",
          start_time: new Date(),
          end_time: new Date(new Date().getTime() + 2 * 60 * 60 * 1000), // 2 hours later
          available_seats: 20,
          createdAt: new Date(),
          updatedAt: new Date(),
        },
      ],
      {}
    );
  },

  async down(queryInterface, _Sequelize) {
    return queryInterface.bulkDelete("Sessions", null, {});
  },
};
