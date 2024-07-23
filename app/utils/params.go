package utils

import "time"

type PlayerQueryParams struct {
  AccelerationRating       uint8     `form:"acceleration_rating" binding:"numeric"`
  Age                      uint8     `form:"age" binding:"numeric"`
  AgilityRating            uint8     `form:"agility_rating" binding:"numeric"`
  AwarenessRating          uint8     `form:"awareness_rating" binding:"numeric"`
  BallCarrierVisionRating  uint8     `form:"ball_carrier_vision_rating" binding:"numeric"`
  BirthDate                time.Time `form:"birth_date"`
  BlockSheddingRating      uint8     `form:"block_shedding_rating" binding:"numeric"`
  BreakSackRating          uint8     `form:"break_sack_rating" binding:"numeric"`
  BreakTackleRating        uint8     `form:"break_tackle_rating" binding:"numeric"`
  CarryingRating           uint8     `form:"carrying_rating" binding:"numeric"`
  CatchInTrafficRating     uint8     `form:"catch_in_traffic_rating" binding:"numeric"`
  CatchingRating           uint8     `form:"catching_rating" binding:"numeric"`
  ChangeOfDirectionRating  uint8     `form:"change_of_direction_rating" binding:"numeric"`
  College                  string    `form:"college"`
  DeepRouteRunningRating   uint8     `form:"deep_route_running_rating" binding:"numeric"`
  FinesseMovesRating       uint8     `form:"finesse_moves_rating" binding:"numeric"`
  FirstName                string    `form:"first_name"`
  Handedness               string    `form:"handedness"`
  Height                   uint8     `form:"height" binding:"numeric"`
  HitPowerRating           uint8     `form:"hit_power_rating" binding:"numeric"`
  ImpactBlockingRating     uint8     `form:"impact_blocking_rating" binding:"numeric"`
  InjuryRating             uint8     `form:"injury_rating" binding:"numeric"`
  JerseyNum                uint8     `form:"jersey_num" binding:"numeric"`
  JukeMoveRating           uint8     `form:"juke_move_rating" binding:"numeric"`
  JumpingRating            uint8     `form:"jumping_rating" binding:"numeric"`
  KickAccuracyRating       uint8     `form:"kick_accuracy_rating" binding:"numeric"`
  KickPowerRating          uint8     `form:"kick_power_rating" binding:"numeric"`
  KickReturnRating         uint8     `form:"kick_return_rating" binding:"numeric"`
  LastName                 string    `form:"last_name"`
  LeadBlockRating          uint8     `form:"lead_block_rating" binding:"numeric"`
  ManCoverageRating        uint8     `form:"man_coverage_rating" binding:"numeric"`
  MediumRouteRunningRating uint8     `form:"medium_route_running_rating" binding:"numeric"`
  OverallRating            uint8     `form:"overall_rating" binding:"numeric"`
  PassBlockFinesseRating   uint8     `form:"pass_block_finesse_rating" binding:"numeric"`
  PassBlockPowerRating     uint8     `form:"pass_block_power_rating" binding:"numeric"`
  PassBlockRating          uint8     `form:"pass_block_rating" binding:"numeric"`
  PlayActionRating         uint8     `form:"play_action_rating" binding:"numeric"`
  PlayRecognitionRating    uint8     `form:"play_recognition_rating" binding:"numeric"`
  Position                 string    `form:"position"`
  PowerMovesRating         uint8     `form:"power_moves_rating" binding:"numeric"`
  PressRating              uint8     `form:"press_rating" binding:"numeric"`
  PursuitRating            uint8     `form:"pursuit_rating" binding:"numeric"`
  ReleaseRating            uint8     `form:"release_rating" binding:"numeric"`
  RunBlockFinesseRating    uint8     `form:"run_block_finesse_rating" binding:"numeric"`
  RunBlockPowerRating      uint8     `form:"run_block_power_rating" binding:"numeric"`
  RunBlockRating           uint8     `form:"run_block_rating" binding:"numeric"`
  RunningStyleRating       string    `form:"running_style_rating"`
  ShortRouteRunningRating  uint8     `form:"short_route_running_rating" binding:"numeric"`
  SpectacularCatchRating   uint8     `form:"spectacular_catch_rating" binding:"numeric"`
  SpeedRating              uint8     `form:"speed_rating" binding:"numeric"`
  SpinMoveRating           uint8     `form:"spin_move_rating" binding:"numeric"`
  StaminaRating            uint8     `form:"stamina_rating" binding:"numeric"`
  StiffArmRating           uint8     `form:"stiff_arm_rating" binding:"numeric"`
  StrengthRating           uint8     `form:"strength_rating" binding:"numeric"`
  TackleRating             uint8     `form:"tackle_rating" binding:"numeric"`
  Team                     string    `form:"team"`
  ThrowAccuracyDeepRating  uint8     `form:"throw_accuracy_deep_rating" binding:"numeric"`
  ThrowAccuracyMidRating   uint8     `form:"throw_accuracy_mid_rating" binding:"numeric"`
  ThrowAccuracyShortRating uint8     `form:"throw_accuracy_short_rating" binding:"numeric"`
  ThrowOnTheRunRating      uint8     `form:"throw_on_the_run_rating" binding:"numeric"`
  ThrowPowerRating         uint8     `form:"throw_power_rating" binding:"numeric"`
  ThrowUnderPressureRating uint8     `form:"throw_under_pressure_rating" binding:"numeric"`
  ToughnessRating          uint8     `form:"toughness_rating" binding:"numeric"`
  TruckingRating           uint8     `form:"trucking_rating" binding:"numeric"`
  Weight                   uint16    `form:"weight" binding:"numeric"`
  NumYears                 uint8     `form:"num_years" binding:"numeric"`
  ZoneCoverageRating       uint8     `form:"zone_coverage_rating" binding:"numeric"`
}
