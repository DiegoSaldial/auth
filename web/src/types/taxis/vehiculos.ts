export interface CreateVehiculos {
  id?: string;
  placa: string;
  puertas: number;
  capacidad: number;
  descripcion: string;
  color: string;
  modelo: string;
  anio: number;
  categoria_id: string;
  foto_url: string;
  caracteristicas: string[];
}

export const CreateVehiculosInput: CreateVehiculos = {
  id: undefined,
  placa: '',
  puertas: 0,
  capacidad: 0,
  descripcion: '',
  color: '',
  modelo: '',
  anio: 0,
  categoria_id: '',
  foto_url: '',
  caracteristicas:[]
}

export interface CreateConductorVehiculos {
  usuario_id: string;
  vehiculo_id: string;
}
