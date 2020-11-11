import React, { ReactNode } from "react";
import Router from "next/router";
import styles from "./Layout.module.css";
import Head from "next/head";
import { Layout as AntLayout, Menu} from "antd";

const { Header, Content, Footer } = AntLayout;

type Props = {
  children?: ReactNode;
  title?: string;
};

const Layout = ({ children, title = "Flow Admin" }: Props) => (
  <>
    <Head>
      <title>{title}</title>
      <meta charSet="utf-8" />
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
    </Head>
    <AntLayout>
      <Header
        style={{
          position: "fixed",
          width: "100%",
          zIndex: 1,
          background: "#fff",
        }}
      >
        <div className={styles.logo} />
        <Menu theme="light" mode="horizontal">
          <Menu.Item
            key="1"
            onClick={() => {
              Router.push("/");
            }}
          >
            Home
          </Menu.Item>
          <Menu.Item
            key="2"
            onClick={() => {
              Router.push("/flow");
            }}
          >
            Flows
          </Menu.Item>
          <Menu.Item
            key="3"
            onClick={() => {
              Router.push("/users");
            }}
          >
            Users
          </Menu.Item>
        </Menu>
      </Header>
      <Content style={{ padding: '0 50px', marginTop: 64 }}>
        <div style={{ padding: 24}}>{children}</div>
      </Content>
      <Footer style={{ textAlign: "center" }}>
        Redshiftkeying ©2020 Created by SlowFlow
      </Footer>
    </AntLayout>
  </>
  // <div>
  //   <Head>
  //     <title>{title}</title>
  //     <meta charSet="utf-8" />
  //     <meta name="viewport" content="initial-scale=1.0, width=device-width" />
  //   </Head>
  //   <header>
  //     <nav>
  //       <Link href="/">
  //         <a>Home</a>
  //       </Link>{" "}
  //       |{" "}
  //       <Link href="/about">
  //         <a>About</a>
  //       </Link>{" "}
  //       |{" "}
  //       <Link href="/users">
  //         <a>Users List</a>
  //       </Link>{" "}
  //       | <a href="/api/users">Users API</a>
  //     </nav>
  //   </header>
  //   {children}
  //   <footer>
  //     <hr />
  //     <span>I'm here to stay (Footer)</span>
  //   </footer>
  // </div>
);

export default Layout;
