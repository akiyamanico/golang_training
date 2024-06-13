import React from "react";
import axios from "axios";

function TambahProduk() {
  const handleSubmit = async (e) => {
    e.preventDefault();

    const formDataToSend = new FormData(e.target);
    
    try {
      const response = await axios.post(
        "http://localhost:8080/produk/tambahproduk",
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
      <h1 className="text-3xl font-bold mb-4">Product Registration</h1>
      <form onSubmit={handleSubmit} encType="multipart/form-data" className="space-y-4">
        <input
          type="text"
          name="nama_produk"
          placeholder="Nama Produk"
          className="border p-2"
          required
        />
        <input
          type="text"
          name="harga_produk"
          placeholder="Harga Produk"
          className="border p-2"
          required
        />
        <input
          type="text"
          name="kategori_produk"
          placeholder="Kategori Produk"
          className="border p-2"
          required
        />
        <input
          type="text"
          name="stok_produk"
          placeholder="Stok Produk"
          className="border p-2"
          required
        />
        <input
          type="file"
          name="gambar_produk"
          className="border p-2"
          required
        />
        <textarea
          name="deskripsi_produk"
          placeholder="Deskripsi Produk"
          className="border p-2"
          required
        />
        <select name="status_produk" className="border p-2" required>
          <option value="">Select Status</option>
          <option value="Available">Available</option>
          <option value="OutOfStock">Out of Stock</option>
        </select>
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

export default TambahProduk;
