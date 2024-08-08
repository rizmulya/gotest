// react component
import { BrowserRouter, Routes, Route } from "react-router-dom";
// layout component
import LandingLayout from "@/layouts/LandingLayout";
import AnotherLayout from "@/layouts/AnotherLayout";
// page component
import Home from "@/pages/Home";
import Users from "@/pages/Users/Index";
import Login from "@/pages/Login";
import Register from "@/pages/Register"

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<LandingLayout />}>
          <Route path="/" element={<Home />}></Route>
          <Route path="/login" element={<Login />}></Route>
          <Route path="/register" element={<Register />}></Route>
        </Route>
        <Route element={<AnotherLayout />}>
          <Route path="/users" element={<Users />}></Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
