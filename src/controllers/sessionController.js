const { Feedback, Session, SessionRegistration, User } = require("../models");
const { Op } = require("sequelize");

exports.getAllSessions = async (_req, res) => {
  try {
    const sessions = await Session.findAll({
      attributes: { exclude: ["createdAt", "updatedAt"] },
    });

    res.status(200).json(sessions);
  } catch (err) {
    res.status(500).json({ message: "Failed to retrieve sessions" });
  }
};

exports.leaveFeedback = async (req, res) => {
  const sessionId = req.params.id;
  const {
    body: { userId, feedback },
  } = req;

  try {
    const session = await Session.findByPk(sessionId);
    if (!session) {
      return res.status(404).json({ message: "Session not found" });
    }

    const user = await User.findByPk(userId);
    if (!user) {
      return res.status(404).json({ message: "User not found" });
    }

    const newFeedback = await Feedback.create({
      session_id: sessionId,
      user_id: userId,
      feedback: feedback,
    });

    res
      .status(201)
      .json({ message: "Feedback successfully subbmitted", newFeedback });
  } catch (err) {
    res.status(500).json({ message: "Failed to leave feedback" });
  }
};

exports.deleteFeedback = async (req, res) => {
  const {
    params: { feedbackId },
    user: { role, id },
  } = req;

  try {
    const feedback = await Feedback.findByPk(feedbackId);

    if (!feedback) {
      return res.status(404).json({ message: "Feedback not found" });
    }

    if (role === "coordinator") {
      await feedback.destroy();
      return res.status(200).json({ message: "Feedback deleted successfully" });
    }

    if (feedback.user_id !== id) {
      return res
        .status(403)
        .json({ message: "You are not authorized to delete this feedback" });
    }

    await feedback.destroy();
    res.status(200).json({ message: "Feedback deleted successfully" });
  } catch (err) {
    res.status(500).json({ message: "Failed to delete feedback" });
  }
};

exports.registerForSession = async (req, res) => {
  const sessionId = req.params.id;
  const {
    body: { userId },
  } = req;

  try {
    const session = await Session.findByPk(sessionId);
    if (!session) {
      return res.status(404).json({ message: "Session not found" });
    }

    if (session.available_seats <= 0) {
      return res.status(400).json({ message: "No seats available" });
    }

    const existingRegistration = await SessionRegistration.findOne({
      where: { user_id: userId, session_id: sessionId },
    });
    if (existingRegistration) {
      return res
        .status(400)
        .json({ message: "Already registered for this session" });
    }

    const overlappingRegistration = await SessionRegistration.findOne({
      where: { user_id: userId },
      include: [
        {
          model: Session,
          where: {
            start_time: { [Op.lte]: session.end_time },
            end_time: { [Op.gte]: session.start_time },
          },
        },
      ],
    });

    if (overlappingRegistration) {
      return res.status(400).json({
        message: "Already registered for another session in this time period",
      });
    }

    await SessionRegistration.create({
      user_id: userId,
      session_id: sessionId,
    });
    session.available_seats -= 1;
    await session.save();

    res.status(201).json({ message: "Registration successful" });
  } catch (err) {
    console.error(err);
    res.status(500).json({ message: "Failed to register for session" });
  }
};

exports.editSession = async (req, res) => {
  const {
    params: { id },
    body: { userId, title, description, startTime, endTime, availableSeats },
  } = req;

  try {
    const session = await Session.findByPk(id);

    if (!session) {
      return res.status(404).json({ message: "Session not found" });
    }

    if (session.user_id !== userId) {
      return res
        .status(403)
        .json({ message: "You are not authorized to edit this session" });
    }

    await session.update({
      title,
      description,
      start_time: startTime,
      end_time: endTime,
      available_seats: availableSeats,
    });
    res.status(200).json({ message: "Session updated successfully" });
  } catch (err) {
    res.status(500).json({ message: "Failed to edit session" });
  }
};

exports.deleteSession = async (req, res) => {
  const sessionId = req.params.id;
  const {
    user: { role, id },
  } = req;

  try {
    const session = await Session.findByPk(sessionId);

    if (!session) {
      return res.status(404).json({ message: "Session not found" });
    }

    if (role === "coordinator") {
      await session.destroy();
      return res.status(200).json({ message: "Session deleted successfully" });
    }

    if (session.user_id !== id) {
      return res
        .status(403)
        .json({ message: "You are not authorized to delete this session" });
    }

    await session.destroy();
    res.status(200).json({ message: "Session deleted successfully" });
  } catch (err) {
    res.status(500).json({ message: "Failed to delete session" });
  }
};
