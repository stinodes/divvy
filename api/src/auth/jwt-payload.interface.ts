export interface JwtPayload {
  nickname: string;
  name: string;
  picture: string;
  updated_at: string;
  email: string;
  email_verified: boolean;
  user_id: string;
}
export interface Auth0User extends JwtPayload {}
