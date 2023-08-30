export interface Viajes {
  id: string;
  pasajero_id: string;
  conductor_id: string;
  estado: boolean;
  descripcion: string;
  categoria_id: string;
  origen_lat: number;
  origen_lon: number;
  destino_lat: number
  destino_lon: number
  registrado: string;
}

export interface CreateViajes {
  pasajero_id: string;
  descripcion: string;
  categoria_id: string;
  origen_lat: number;
  origen_lon: number;
  destino_lat: number;
  destino_lon: number;
}
export const CrearInput: CreateViajes = {
  pasajero_id: '',
  descripcion: '',
  categoria_id: '',
  origen_lat: 0,
  origen_lon: 0,
  destino_lat: 0,
  destino_lon: 0,
}

export interface QueryViajes {
  fecha_inicio: string;
  fecha_fin: string;
  include_disponibles: boolean;
}
export const QueryInput: QueryViajes = {
  fecha_inicio: '',
  fecha_fin: '',
  include_disponibles: true,
}
