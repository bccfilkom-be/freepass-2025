const { User } = require("../models");

exports.userValidationSchema = {
  username: {
    in: ["body"],
    isEmpty: {
      negated: true,
      errorMessage: "Username is required",
    },
    isString: {
      errorMessage: "Username must be a string",
    },
  },
  email: {
    in: ["body"],
    isEmail: {
      errorMessage: "Invalid email format",
    },
    custom: {
      options: async (email) => {
        const existingUser = await User.findOne({ where: { email } });
        if (existingUser) {
          throw new Error("Email already in use");
        }
      },
    },
  },
  password: {
    in: ["body"],
    isLength: {
      options: { min: 8 },
      errorMessage: "Password must be at least 8 characters long",
    },
    matches: {
      options: [/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/],
      errorMessage:
        "Password must contain at least one uppercase letter, one lowercase letter, and one number",
    },
  },
};

exports.updateUserValidationSchema = {
  id: {
    in: ["params"],
    isUUID: {
      errorMessage: "Invalid user ID",
    },
  },
  username: {
    in: ["body"],
    optional: true,
    isEmpty: {
      negated: true,
      errorMessage: "Username is required",
    },
    isString: {
      errorMessage: "Username must be a string",
    },
  },
  email: {
    in: ["body"],
    optional: true,
    isEmail: {
      errorMessage: "Invalid email format",
    },
    custom: {
      options: async (email) => {
        const existingUser = await User.findOne({ where: { email } });
        if (existingUser) {
          throw new Error("Email already in use");
        }
      },
    },
  },
  password: {
    in: ["body"],
    optional: true,
    isLength: {
      options: { min: 8 },
      errorMessage: "Password must be at least 8 characters long",
    },
    matches: {
      options: [/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/],
      errorMessage:
        "Password must contain at least one uppercase letter, one lowercase letter, and one number",
    },
  },
};
