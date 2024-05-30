import { useState, useEffect, useMemo } from "react";
import { Container, Box, Typography, Button, Paper } from "@mui/material";
import { DataGrid, GridCellEditStopReasons } from "@mui/x-data-grid";
import AddIcon from "@mui/icons-material/Add";
import moment from "moment";
import UserActions from "./UserActions";
import { fetchUsers } from "../api/axios";
import { Link } from "react-router-dom";

export default function Users() {
  const [rows, setRows] = useState([]);
  const [rowId, setRowId] = useState(null);

  useEffect(() => {
    const fetchUsersData = async () => {
      const userData = await fetchUsers();
      setRows(userData.data);
    };

    fetchUsersData();
  }, []);

  const columns = useMemo(
    () => [
      { field: "id", headerName: "ID", width: 60 },
      {
        field: "fullName",
        headerName: "Full name",
        description: "This column has a value getter and is not sortable.",
        sortable: false,
        width: 200,
        valueGetter: (params) => `${params.row.fullName}`,
      },
      {
        field: "email",
        headerName: "Email",
        width: 250,
        editable: true,
      },
      {
        field: "roleId",
        headerName: "Role",
        type: "singleSelect",
        valueGetter: (params) => {
          switch (params.row.roleId) {
            case 1:
              return "Ophthalmologist";
            case 2:
              return "Nurse";
            case 3:
              return "Technician";
            default:
              return params.row.roleId;
          }
        },
        valueOptions: ["Ophthalmologist", "Nurse", "Technician"],
        width: 150,
        editable: true,
      },
      {
        field: "createdAt",
        headerName: "Create At",
        renderCell: (params) =>
          moment(params.row.createdAt).format("YYYY-MM-DD HH:MM:SS"),
        width: 200,
      },
      {
        field: "active",
        headerName: "Active",
        type: "boolean",
        width: 100,
        editable: true,
      },
      {
        field: "actions",
        headerName: "Actions",
        type: "actions",
        renderCell: (params) => (
          <UserActions {...{ params, rowId, setRowId }} />
        ),
      },
    ],
    [rowId]
  );

  return (
    <Container maxWidth="lg" sx={{ mt: 3, mb: 3 }}>
      <Box
        display="flex"
        alignItems="end"
        justifyContent="space-between"
        sx={{ mb: 1 }}
      >
        <Typography variant="h5" component="h5">
          Manage Users
        </Typography>

        <Link to="register">
          <Button
            variant="contained"
            endIcon={<AddIcon />}
            sx={{ textTransform: "none" }}
          >
            Create User
          </Button>
        </Link>
      </Box>

      <DataGrid
        rows={rows}
        columns={columns}
        initialState={{
          pagination: {
            paginationModel: {
              pageSize: 5,
            },
          },
        }}
        pageSizeOptions={[5, 10]}
        onCellEditStart={(params) => {
          setRowId(params.id);
        }}
       
      />
    </Container>
  );
}
