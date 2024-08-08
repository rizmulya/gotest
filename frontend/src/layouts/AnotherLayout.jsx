import Header from "@/components/Header";
import { Outlet } from "react-router-dom";

const AnotherLayout = () => {
  return (
    <>
      layout: AnotherLayout
      <Header />
      <form action="/logout" method="post">
        <button>Logout</button>
      </form>
      <Outlet />
    </>
  );
};

export default AnotherLayout;
