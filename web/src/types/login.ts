export interface LoginData {
  username: string;
  password: string;
}

export interface ResponseLogin {
  id: string;
  nombre: string;
  bit: number;
}

export const Instancia: LoginData = {
  username: '',
  password: '',
};
