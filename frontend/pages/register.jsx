import Header from "../components/Layouts/Header";
import MainLayout from "../components/Layouts/MainLayout";
import Layout from "../components/Auth/Layout";
import Banner from "../components/Auth/Banner";
import RegisterForm from "../components/Auth/RegisterForm";
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
          <RegisterForm></RegisterForm>
          <Banner></Banner>
        </Layout>
      </MainLayout>
    </>
  );
}
