import { useState } from "react";
import {
  Box,
  Container,
  TextField,
  FormControlLabel,
  FormControl,
  Radio,
  RadioGroup,
  FormLabel,
  Button,
  Typography,
  Paper,
  InputAdornment,
  IconButton,
} from "@mui/material";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import Grid2 from "@mui/material/Unstable_Grid2";
import { createUser } from "../api/axios";
import { Link, useNavigate } from "react-router-dom";

const isValidEmail = (email) => {
  const emailRegex = /^([a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$/;
  return emailRegex.test(email);
};

const isValidPassword = (password) => {
  const passwordRegex = /^(?=.*[a-z])(?=.*\d)[a-zA-Z\d]{8,}$/;
  return passwordRegex.test(password);
};

const isValidName = (name) => {
  const nameRegex = /^[^0-9]+$/;
  return nameRegex.test(name);
};

export default function RegisterForm() {
  const [roleId, setRoleId] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [emailError, setEmailError] = useState(false);
  const [passwordError, setPasswordError] = useState(false);
  const [firstNameError, setFirstNameError] = useState(false);
  const [lastNameError, setLastNameError] = useState(false);
  const [showPassword, setShowPassword] = useState(false);

  const navigate = useNavigate();

  const handleClickShowPassword = () => setShowPassword((show) => !show);

  const handleMouseDownPassword = (e) => {
    e.preventDefault();
  };

  const preventCopy = (e) => {
    e.preventDefault();
    e.stopPropagation();
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const isEmailValid = isValidEmail(email);
    const isPasswordValid = isValidPassword(password);
    const isFirstNameValid = isValidName(firstName);
    const isLastNameValid = isValidName(lastName);

    setEmailError(!isEmailValid);
    setPasswordError(!isPasswordValid);
    setFirstNameError(!isFirstNameValid);
    setLastNameError(!isLastNameValid);

    if (
      isEmailValid &&
      isPasswordValid &&
      isFirstNameValid &&
      isLastNameValid &&
      password === confirmPassword
    ) {
      const user = {
        Firstname: firstName,
        Lastname: lastName,
        Email: email,
        Password: password,
        RoleId: parseInt(roleId),
      };

      await createUser(user);
      alert(
        "The account is created, Please wait for the admin to approve the account."
      );
      navigate('/');
    }
  };

  return (
    <Box
      display="flex"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
      margin={2}
    >
      <Container maxWidth="xs">
        <Link to="/">
          <Button variant="contained" startIcon={<ArrowBackIcon />}>
            Back
          </Button>
        </Link>

        <Paper elevation={3}>
          <Box
            onSubmit={handleSubmit}
            component="form"
            sx={{ borderRadius: 2, mt: 2, padding: 2 }}
          >
            <Typography
              component="h1"
              variant="h4"
              align="center"
              sx={{ padding: 2 }}
            >
              Create Account
            </Typography>

            <Grid2 container spacing={2}>
              <Grid2 item md={12}>
                <FormControl required>
                  <FormLabel style={{ textAlign: "left" }}>Role</FormLabel>
                  <RadioGroup
                    row
                    id="roleId"
                    name="roleId"
                    aria-labelledby="roleId"
                    onChange={(e) => {
                      setRoleId(e.target.value);
                    }}
                    value={roleId}
                  >
                    <FormControlLabel
                      value={1}
                      control={<Radio />}
                      label="Doctor"
                    />
                    <FormControlLabel
                      value={2}
                      control={<Radio />}
                      label="Nurse"
                    />
                    <FormControlLabel
                      value={3}
                      control={<Radio />}
                      label="Technician"
                    />
                  </RadioGroup>
                </FormControl>
              </Grid2>

              <Grid2 item xs={12} md={6}>
                <TextField
                  required
                  fullWidth
                  id="firstName"
                  name="firstName"
                  label="First Name"
                  onChange={(e) => {
                    setFirstName(e.target.value);
                    setFirstNameError(!isValidName(e.target.value));
                  }}
                  value={firstName}
                  error={firstNameError}
                  helperText={firstNameError ? "Please enter only letters" : ""}
                  inputProps={{ maxLength: 100 }}
                />
              </Grid2>

              <Grid2 item xs={12} md={6}>
                <TextField
                  required
                  fullWidth
                  id="lastName"
                  name="lastName"
                  label="Last Name"
                  onChange={(e) => {
                    setLastName(e.target.value);
                    setLastNameError(!isValidName(e.target.value));
                  }}
                  value={lastName}
                  error={lastNameError}
                  helperText={lastNameError ? "Please enter only letters" : ""}
                  inputProps={{ maxLength: 100 }}
                />
              </Grid2>

              <Grid2 item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="email"
                  name="email"
                  label="Email"
                  onChange={(e) => {
                    setEmail(e.target.value);
                    setEmailError(!isValidEmail(e.target.value));
                  }}
                  value={email}
                  error={emailError}
                  helperText={
                    emailError ? "Please enter the correct email format." : ""
                  }
                  inputProps={{ maxLength: 255 }}
                />
              </Grid2>

              <Grid2 item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="password"
                  name="password"
                  label="Password"
                  type={showPassword ? "text" : "password"}
                  onCopy={preventCopy}
                  onCut={preventCopy}
                  onPaste={(e) => {
                    e.preventDefault();
                    setPassword("");
                  }}
                  onChange={(e) => {
                    setPassword(e.target.value);
                    setPasswordError(!isValidPassword(e.target.value));
                  }}
                  inputProps={{ maxLength: 255 }}
                  value={password}
                  error={passwordError}
                  helperText={
                    passwordError
                      ? "Password must contain at least 8 lowercase letters and numbers."
                      : ""
                  }
                  InputProps={{
                    endAdornment: (
                      <InputAdornment position="end">
                        <IconButton
                          aria-label="toggle password visibility"
                          onClick={handleClickShowPassword}
                          onMouseDown={handleMouseDownPassword}
                          edge="end"
                        >
                          {showPassword ? <Visibility /> : <VisibilityOff />}
                        </IconButton>
                      </InputAdornment>
                    ),
                  }}
                />
              </Grid2>

              <Grid2 item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="confirmPassword"
                  name="confirmPassword"
                  label="Confirm Password"
                  type="password"
                  onCopy={preventCopy}
                  onCut={preventCopy}
                  onPaste={(e) => {
                    e.preventDefault();
                    setConfirmPassword("");
                  }}
                  onChange={(e) => {
                    setConfirmPassword(e.target.value);
                  }}
                  value={confirmPassword}
                  error={confirmPassword !== password}
                  helperText={
                    confirmPassword !== password
                      ? "Please make sure your password match."
                      : ""
                  }
                  inputProps={{ maxLength: 255 }}
                />
              </Grid2>
            </Grid2>

            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ fontSize: 20, textTransform: "none", mt: 3 }}
            >
              Sign up
            </Button>
          </Box>
        </Paper>
      </Container>
    </Box>
  );
}
