import { Tab, TabList } from '@fluentui/react-components';
import { Link } from '@tanstack/react-router';

export default function Nav() {
  return (
    <>
      <nav className="flex justify-center items-center mb-5">
        <h2 className="text-xl font-bold">MallData</h2>
        <TabList defaultSelectedValue={location.href.includes('map') ? 'map' : 'index'}>
          <Link to="/">
            <Tab value="index">Shops</Tab>
          </Link>
          <Link to="/map">
            <Tab value="map">Map</Tab>
          </Link>
          <Link to="/dashboard">
            <Tab value="dashboard">Dashboard</Tab>
          </Link>
        </TabList>
      </nav>
    </>
  );
}
