import {
   Injectable,
   CanActivate,
   ExecutionContext,
   UnauthorizedException,
} from "@nestjs/common"
import { JwtService } from "@nestjs/jwt"
import { Request } from "express"
import { CookieService } from "./cookie.service"
import { GetSessionInfoDto } from "./dto"

@Injectable()
export class AuthGuard implements CanActivate {
   constructor(private jwtService: JwtService) {}

   async canActivate(context: ExecutionContext) {
      const req = context.switchToHttp().getRequest() as Request
      const token = req.cookies && req.cookies[CookieService.tokenKey]

      if (!token) {
         throw new UnauthorizedException()
      }

      try {
         const sessionInfo: GetSessionInfoDto = await this.jwtService.verifyAsync(token, {
            secret: process.env.JWT_SECRET || "secret_key",
         })

         req["session"] = sessionInfo
      } catch (error) {
         console.error("JWT verification error:", error)
         throw new UnauthorizedException()
      }

      return true
   }
}
