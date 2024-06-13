import React from "react";
import axios from "axios";

function App() {
  const handleSubmit = async (e) => {
    e.preventDefault();

    const formDataToSend = new FormData(e.target);
    
    try {
      const response = await axios.post(
        "http://localhost:8080/usersCreate",
        formDataToSend,
        {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        }
      );
      console.log(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="container mx-auto">
      <h1 className="text-3xl font-bold mb-4">User Registration</h1>
      <form onSubmit={handleSubmit} encType="multipart/form-data" className="space-y-4">
        <input
          type="text"
          name="username"
          placeholder="Username"
          className="border p-2"
          required
        />
        <input
          type="email"
          name="email"
          placeholder="Email"
          className="border p-2"
          required
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          className="border p-2"
          required
        />
        <input
          type="text"
          name="nama"
          placeholder="Nama"
          className="border p-2"
          required
        />
        <input
          type="tel"
          name="no_telp"
          placeholder="Nomor Telepon"
          className="border p-2"
          required
        />
        <input
          type="date"
          name="tanggal_lahir"
          className="border p-2"
          required
        />
        <select name="jenis_kelamin" className="border p-2" required>
          <option value="">Select Gender</option>
          <option value="Male">Male</option>
          <option value="Female">Female</option>
        </select>
        <input
          type="file"
          name="profile_picture"
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

export default App;
