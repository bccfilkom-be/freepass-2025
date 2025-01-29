const createSessionValidationSchema = (isUpdate = false) => {
  return {
    id: {
      in: ["params"],
      isInt: {
        errorMessage: "Session ID must be an integer",
      },
    },
    userId: {
      in: ["body"],
      isUUID: {
        errorMessage: "User ID must be a UUID",
      },
      optional: isUpdate,
    },
    title: {
      in: ["body"],
      isString: {
        errorMessage: "Title must be a string",
      },
      notEmpty: {
        errorMessage: "Title is required",
      },
      optional: isUpdate,
    },
    description: {
      in: ["body"],
      isString: {
        errorMessage: "Description must be a string",
      },
      notEmpty: {
        errorMessage: "Description is required",
      },
      optional: isUpdate,
    },
    startTime: {
      in: ["body"],
      isISO8601: {
        errorMessage: "Start time must be a valid date",
      },
      optional: isUpdate,
    },
    endTime: {
      in: ["body"],
      isISO8601: {
        errorMessage: "End time must be a valid date",
      },
      optional: isUpdate,
    },
    availableSeats: {
      in: ["body"],
      isInt: {
        options: { min: 1 },
        errorMessage: "Seats must be at least 1",
      },
      optional: isUpdate,
    },
  };
};

const leaveFeedbackSchema = {
  id: {
    in: ["params"],
    isInt: {
      errorMessage: "Session ID must be an integer",
    },
  },
  feedback: {
    in: ["body"],
    isString: {
      errorMessage: "Feedback must be a string",
    },
    notEmpty: {
      errorMessage: "Feedback is required",
    },
  },
  userId: {
    in: ["body"],
    isUUID: {
      errorMessage: "User ID must be a UUID",
    },
  },
};

const registerForSessionSchema = {
  id: {
    in: ["params"],
    isInt: {
      errorMessage: "Session ID must be an integer",
    },
  },
  userId: {
    in: ["body"],
    isUUID: {
      errorMessage: "User ID must be a UUID",
    },
  },
};

module.exports = {
  createSessionValidationSchema,
  leaveFeedbackSchema,
  registerForSessionSchema,
};
