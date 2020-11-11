import Link from "next/link";
import Layout from "../components/Layout";

const IndexPage = () => (
  <Layout title="Home">
    <h1>Hello Slow Flow</h1>
    <p>
      <Link href="/about">
        <a>About</a>
      </Link>
    </p>
    <p>
      <Link href="/flow/flows">
        <a>工作流</a>
      </Link>
    </p>
  </Layout>
);

export default IndexPage;
