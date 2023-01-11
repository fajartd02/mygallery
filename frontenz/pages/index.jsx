import Header from "../components/Layouts/Header";
import HeroSection from "../components/Home/HeroSection";
import BenefitSection from "../components/Home/BenefitSection";
import SignInSection from "../components/Home/SignInSection";
import Footer from "../components/Home/Footer";

export default function Home() {
  return (
    <>
      <Header />
      <HeroSection />
      <BenefitSection />
      <SignInSection />
      <Footer />
    </>
  );
}
