import React, { useEffect, useState } from "react";
import { Box, Button, CircularProgress, Fab } from "@mui/material";
import Check from "@mui/icons-material/Check";
import Save from "@mui/icons-material/Save";
import { green } from "@mui/material/colors";
import { updateUserPartial } from "../api/axios";

export default function UserActions({ params, rowId, setRowId }) {
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);

  const handleSubmit = async () => {
    const { id, email, active } = params.row;
    let { roleId } = params.row;
    switch (roleId) {
      case "Ophthalmologist":
        roleId = 1;
        break;
      case "Nurse":
        roleId = 2;
        break;
      case "Technician":
        roleId = 3;
        break;
    }

    const user = {
      Email: email,
      RoleId: roleId,
      Active: active
    }

    setLoading(true);
    const result = await updateUserPartial(user, id);
    if (result) {
      setSuccess(true);
      setRowId(null);
    }
    setLoading(false);
    setTimeout(() => {
      setSuccess(false);
    }, 1500)
  };

  useEffect(() => {
    if (rowId === params.id && success) {
      setSuccess(false);
    }
  }, [rowId]);

  return (
    <Box sx={{ m: 1, position: "relative" }}>
      {success ? (
        <Fab
          color="primary"
          sx={{
            width: 40,
            height: 40,
            bgcolor: green[500],
            "&:hover": { bgcolor: green[700] },
          }}
        >
          <Check />
        </Fab>
      ) : (
        <Fab
          color="primary"
          sx={{ width: 40, height: 40 }}
          disabled={params.id !== rowId || loading}
          onClick={handleSubmit}
        >
          <Save />
        </Fab>
      )}
      {loading && (
        <CircularProgress
          size={52}
          sx={{
            color: green[500],
            position: "absolute",
            top: -6,
            left: -6,
            zIndex: 1,
          }}
        />
      )}
    </Box>
  );
}
