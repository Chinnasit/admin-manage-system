import axios from "axios";

const API_BASE_URL = "http://localhost:3000";

export const fetchUsers = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/users`);
    return response;
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
};

export const createUser = async (user) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/user`, user);
    return response;
  } catch (error) {
    console.error("Error creating user:", error);
    throw error;
  }
};

export const updateUserPartial = async (user, id) => {
    try {
        const response = await axios.patch(`${API_BASE_URL}/users/${id}`, user)
        return response
    } catch (error) {
        console.error("Error updating user:", error);
    throw error;
    }
}