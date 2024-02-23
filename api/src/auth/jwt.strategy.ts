import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { PassportStrategy } from '@nestjs/passport';
import * as dotenv from 'dotenv';
import { passportJwtSecret } from 'jwks-rsa';
import { ExtractJwt, Strategy } from 'passport-jwt';
import { JwtPayload } from './jwt-payload.interface';

dotenv.config();

@Injectable()
export class JWTStrategy extends PassportStrategy(Strategy) {
  constructor(private configService: ConfigService) {
    const config = {
      secretOrKeyProvider: passportJwtSecret({
        cache: true,
        rateLimit: true,
        jwksRequestsPerMinute: 5,
        jwksUri: `${configService.get('AUTH0_URL')}/.well-known/jwks.json`,
      }),

      jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
      aud: `${configService.get('AUTH0_AUDIENCE')}`,
      issuer: `${configService.get('AUTH0_URL')}/`,
      algorithms: ['RS256'],
    };
    super(config);
  }

  validate(payload: JwtPayload): JwtPayload {
    return payload;
  }
}
