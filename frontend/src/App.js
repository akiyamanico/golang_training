import TambahProduk from "./views/Produk/TambahProduk";
import UserRegistration from "./views/User/UserRegistration";
import { BrowserRouter, Routes, Route } from "react-router-dom";
function app() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<UserRegistration />}></Route>
      </Routes>
      
      <Routes>
        <Route exact path="/tambahproduk" element={<TambahProduk />}></Route>
      </Routes>
      
    </BrowserRouter>
  );
}

export default app;
