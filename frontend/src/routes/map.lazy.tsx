import { createLazyFileRoute } from '@tanstack/react-router';
import { MapContainer, TileLayer } from 'react-leaflet';

export const Route = createLazyFileRoute('/map')({
  component: MapView,
});

function MapView() {
  return (
    <div className="w-full h-full">
      <link
        rel="stylesheet"
        href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
        integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
      />
      <script
        src="https://unpkg.com/leaflet@1.9.4/dist/leaflet.js"
        integrity="sha256-20nQCchB9co0qIjJZRGuk2/Z9VM+kNiyxNV1lvTlZBo="></script>

      <MapContainer>
        <TileLayer attribution='<a href="https://sc3.io">SwitchCraft3</a>' url="" />
      </MapContainer>
    </div>
  );
}
