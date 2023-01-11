import Header from "../components/Layouts/Header";
import MainLayout from "../components/Layouts/MainLayout";
import Layout from "../components/Auth/Layout";
import LoginForm from "../components/Auth/LoginForm";
import Banner from "../components/Auth/Banner";
import { useEffect } from "react";
import { redirectIfAuthenticated } from "../helper/auth";

export default function Login() {
  useEffect(() => {
    redirectIfAuthenticated();
  }, []);

  return (
    <>
      <Header />
      <MainLayout>
        <Layout>
          <LoginForm></LoginForm>
          <Banner></Banner>
        </Layout>
      </MainLayout>
    </>
  );
}
