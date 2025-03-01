import { StoryService } from "./story.service"
import { OrderByFilter, SortByScenesAmount } from "./types/types"
import { Response } from "express"
import {
   Controller,
   Get,
   Res,
   Param,
   NotFoundException,
   Patch,
   Query,
   UseGuards,
   UseInterceptors,
   Body,
   Put,
} from "@nestjs/common"
import { AuthGuard } from "src/auth/auth.guard"
import { GetSessionInfoDto } from "src/auth/dto"
import { SessionInfo } from "src/auth/sessionInfoDecorator"
import { SessionInterceptor } from "src/auth/sessionInterseptor"
import { CreateStoryResultDto } from "./story.dto"
import { SceneService } from "src/scene/scene.service"

@Controller("stories")
export class StoryController {
   constructor(
      private storyService: StoryService,
      private sceneService: SceneService,
   ) {}

   @UseInterceptors(SessionInterceptor)
   @Get()
   async getStories(
      @Res() res: Response,
      @SessionInfo() session: GetSessionInfoDto,
      @Query("limit") limit?: number,
      @Query("page") page?: number,
      @Query("search") search: string = "",
      @Query("length") length: SortByScenesAmount = "",
      @Query("order") order: OrderByFilter = "",
      @Query("by_user") byUser?: string,
      @Query("only_count") onlyCount: boolean = false,
   ) {
      const userId = session?.id
      const count = await this.storyService.getStoryCount(search, length)

      res.setHeader("X-Total-Count", count)
      res.setHeader("Access-Control-Expose-Headers", "X-Total-Count")

      if (!onlyCount) {
         const stories = await this.storyService.getAllStories(
            search,
            length,
            order,
            userId,
            byUser,
            limit,
            page,
         )

         return res.json(stories)
      }

      return res.json([])
   }

   @UseInterceptors(SessionInterceptor)
   @Get(":storyId")
   async getStoryById(
      @Param("storyId") storyId: string,
      @SessionInfo() session: GetSessionInfoDto,
   ) {
      const userId = session?.id
      const story = await this.storyService.getStoryById(storyId, userId)

      if (!story) {
         throw new NotFoundException()
      }

      return story
   }

   @Patch(":storyId/passes")
   async updatePasses(@Param("storyId") storyId: string) {
      const updatedPasses = await this.storyService.updatePasses(storyId)
      return updatedPasses
   }

   @Patch(":storyId/like")
   @UseGuards(AuthGuard)
   async toggleLike(
      @Param("storyId") storyId: string,
      @SessionInfo() session: GetSessionInfoDto,
   ) {
      const userId = session.id
      const res = await this.storyService.toggleLike(storyId, userId)
      return res
   }

   @Get(":storyId/results/:userId")
   async getUserResult(
      @Param("storyId") storyId: string,
      @Param("userId") userId: string,
   ) {
      const res = await this.storyService.getUserResult({ storyId, userId })

      if (!res) {
         throw new NotFoundException()
      }

      const scene = await this.sceneService.getScene({
         storyId,
         number: res.resultSceneNumber,
      })

      return { ...res, scene }
   }

   @Put(":storyId/results/:userId")
   @UseGuards(AuthGuard)
   async setResult(
      @Body() body: CreateStoryResultDto,
      @Param("storyId") storyId: string,
      @Param("userId") userId: string,
   ) {
      const res = await this.storyService.setResult({ storyId, userId, ...body })
      return res
   }

   @Get(":storyId/scenes")
   async getScenesByStoryId(@Param("storyId") storyId: string) {
      return await this.sceneService.getScenesByStoryId(storyId)
   }

   @Get(":storyId/scenes/:sceneNumber")
   async getScene(
      @Param("storyId") storyId: string,
      @Param("sceneNumber") sceneNumber: string,
   ) {
      const res = await this.sceneService.getScene({ storyId, number: sceneNumber })
      return res
   }

   @Patch(":storyId/passes/:sceneId")
   async incrementPasses(
      @Param("storyId") storyId: string,
      @Param("sceneId") sceneId: string,
   ) {
      return await this.sceneService.incrementPasses(storyId, sceneId)
   }

   @Get(":storyId/results")
   async getStatistics(@Param("storyId") storyId: string) {
      const endScenes = await this.sceneService.getEndScenes(storyId)
      return endScenes
   }
}
