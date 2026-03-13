const express = require("express");
const asyncHandler = require("express-async-handler");
const { getStudentDetail } = require("./students-service");

const router = express.Router();

router.get("/:id", asyncHandler(async (req, res) => {
    const { id } = req.params;
    const student = await getStudentDetail(id);
    res.json(student);
}));

module.exports = { studentsInternalRoutes: router };
