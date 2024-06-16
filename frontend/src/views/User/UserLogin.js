import React, { useState } from "react";
import axios from "axios";

function UserLogin() {
    const [credentials, setCredentials] = useState({
        username : '',
        password : ''
    });
    const handleChange = (e) => {
        setCredentials({
          ...credentials,
          [e.target.name]: e.target.value
        });
      };
  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post(
        "http://localhost:8080/users/usersLogin",
        credentials,
      );
      const {tokenString} = response.data;
      localStorage.setItem('token', tokenString);
      console.log(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="container mx-auto">
      <h1 className="text-3xl font-bold mb-4">User Registration</h1>
      <form onSubmit={handleSubmit}  className="space-y-4">
        <input
          type="text"
          name="username"
          value={credentials.usernamel}
          onChange={handleChange}
          placeholder="Username"
          className="border p-2"
          required
        />
        <input
          type="password"
          name="password"
          value={credentials.password}
          onChange={handleChange}
          placeholder="Password"
          className="border p-2"
          required
        />
        <button
          type="submit"
          className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
        >
          Submit
        </button>
      </form>
    </div>
  );
}

export default UserLogin;
