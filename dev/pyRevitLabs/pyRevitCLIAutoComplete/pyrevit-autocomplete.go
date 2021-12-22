package main

import (
	"github.com/posener/complete"
)

func main() {
	pyrevit := complete.Command{
		Sub: complete.Commands{
			"wiki": complete.Command{
				Sub:   complete.Commands{},
				Flags: complete.Flags{},
			},
			"blog": complete.Command{
				Sub:   complete.Commands{},
				Flags: complete.Flags{},
			},
			"docs": complete.Command{
				Sub:   complete.Commands{},
				Flags: complete.Flags{},
			},
			"source": complete.Command{
				Sub:   complete.Commands{},
				Flags: complete.Flags{},
			},
			"youtube": complete.Command{
				Sub:   complete.Commands{},
				Flags: complete.Flags{},
			},
			"support": complete.Command{
				Sub:   complete.Commands{},
				Flags: complete.Flags{},
			},
			"env": complete.Command{
				Sub: complete.Commands{},
				Flags: complete.Flags{
					"--log":  complete.PredictNothing,
					"--json": complete.PredictNothing,
					"--help": complete.PredictNothing,
				},
			},
			"update": complete.Command{
				Sub: complete.Commands{},
				Flags: complete.Flags{
					"--help": complete.PredictNothing,
				},
			},
			"clone": complete.Command{
				Sub: complete.Commands{},
				Flags: complete.Flags{
					"--image":    complete.PredictNothing,
					"--password": complete.PredictNothing,
					"--help":     complete.PredictNothing,
					"--token":    complete.PredictNothing,
					"--branch":   complete.PredictNothing,
					"--log":      complete.PredictNothing,
					"--dest":     complete.PredictNothing,
				},
			},
			"clones": complete.Command{
				Sub: complete.Commands{
					"info": complete.Command{
						Sub:   complete.Commands{},
						Flags: complete.Flags{},
					},
					"open": complete.Command{
						Sub:   complete.Commands{},
						Flags: complete.Flags{},
					},
					"add": complete.Command{
						Sub: complete.Commands{
							"this": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--force": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--force": complete.PredictNothing,
						},
					},
					"forget": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"rename": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"delete": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--clearconfigs": complete.PredictNothing,
						},
					},
					"branch": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"version": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"commit": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"origin": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log":   complete.PredictNothing,
							"--reset": complete.PredictNothing,
						},
					},
					"update": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--token":    complete.PredictNothing,
							"--log":      complete.PredictNothing,
							"--password": complete.PredictNothing,
						},
					},
					"deployments": complete.Command{
						Sub:   complete.Commands{},
						Flags: complete.Flags{},
					},
					"engines": complete.Command{
						Sub:   complete.Commands{},
						Flags: complete.Flags{},
					},
				},
				Flags: complete.Flags{
					"--help": complete.PredictNothing,
				},
			},
			"attach": complete.Command{
				Sub: complete.Commands{
					"default": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--attached":  complete.PredictNothing,
							"--allusers":  complete.PredictNothing,
							"--installed": complete.PredictNothing,
						},
					},
				},
				Flags: complete.Flags{
					"--attached":  complete.PredictNothing,
					"--help":      complete.PredictNothing,
					"--allusers":  complete.PredictNothing,
					"--installed": complete.PredictNothing,
				},
			},
			"attached": complete.Command{
				Sub: complete.Commands{},
				Flags: complete.Flags{
					"--help": complete.PredictNothing,
				},
			},
			"switch": complete.Command{
				Sub: complete.Commands{},
				Flags: complete.Flags{
					"--help": complete.PredictNothing,
				},
			},
			"detach": complete.Command{
				Sub: complete.Commands{},
				Flags: complete.Flags{
					"--log":  complete.PredictNothing,
					"--help": complete.PredictNothing,
				},
			},
			"extend": complete.Command{
				Sub: complete.Commands{
					"ui": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--token":    complete.PredictNothing,
							"--log":      complete.PredictNothing,
							"--password": complete.PredictNothing,
							"--dest":     complete.PredictNothing,
						},
					},
					"lib": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--token":    complete.PredictNothing,
							"--log":      complete.PredictNothing,
							"--password": complete.PredictNothing,
							"--dest":     complete.PredictNothing,
						},
					},
				},
				Flags: complete.Flags{
					"--password": complete.PredictNothing,
					"--help":     complete.PredictNothing,
					"--token":    complete.PredictNothing,
					"--log":      complete.PredictNothing,
					"--dest":     complete.PredictNothing,
				},
			},
			"extensions": complete.Command{
				Sub: complete.Commands{
					"search": complete.Command{
						Sub:   complete.Commands{},
						Flags: complete.Flags{},
					},
					"info": complete.Command{
						Sub:   complete.Commands{},
						Flags: complete.Flags{},
					},
					"help": complete.Command{
						Sub:   complete.Commands{},
						Flags: complete.Flags{},
					},
					"open": complete.Command{
						Sub:   complete.Commands{},
						Flags: complete.Flags{},
					},
					"delete": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"origin": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log":   complete.PredictNothing,
							"--reset": complete.PredictNothing,
						},
					},
					"paths": complete.Command{
						Sub: complete.Commands{
							"forget": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--all": complete.PredictNothing,
									"--log": complete.PredictNothing,
								},
							},
							"add": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log":  complete.PredictNothing,
							"--help": complete.PredictNothing,
						},
					},
					"enable": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"disable": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"sources": complete.Command{
						Sub: complete.Commands{
							"forget": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--all": complete.PredictNothing,
									"--log": complete.PredictNothing,
								},
							},
							"add": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log":  complete.PredictNothing,
							"--help": complete.PredictNothing,
						},
					},
					"update": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--token":    complete.PredictNothing,
							"--log":      complete.PredictNothing,
							"--password": complete.PredictNothing,
						},
					},
				},
				Flags: complete.Flags{
					"--log":  complete.PredictNothing,
					"--help": complete.PredictNothing,
				},
			},
			"releases": complete.Command{
				Sub: complete.Commands{
					"latest": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--pre": complete.PredictNothing,
						},
					},
					"open": complete.Command{
						Sub: complete.Commands{
							"latest": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--pre": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--pre": complete.PredictNothing,
						},
					},
					"download": complete.Command{
						Sub: complete.Commands{
							"installer": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--dest": complete.PredictNothing,
								},
							},
							"archive": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--dest": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--dest": complete.PredictNothing,
						},
					},
				},
				Flags: complete.Flags{
					"--pre":  complete.PredictNothing,
					"--help": complete.PredictNothing,
				},
			},
			"revits": complete.Command{
				Sub: complete.Commands{
					"killall": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"fileinfo": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--rft": complete.PredictNothing,
							"--csv": complete.PredictNothing,
							"--rte": complete.PredictNothing,
						},
					},
				},
				Flags: complete.Flags{
					"--supported": complete.PredictNothing,
					"--help":      complete.PredictNothing,
					"--installed": complete.PredictNothing,
				},
			},
			"run": complete.Command{
				Sub: complete.Commands{
					"commands": complete.Command{
						Sub:   complete.Commands{},
						Flags: complete.Flags{},
					},
				},
				Flags: complete.Flags{
					"--import":       complete.PredictNothing,
					"--purge":        complete.PredictNothing,
					"--help":         complete.PredictNothing,
					"--revit":        complete.PredictNothing,
					"--models":       complete.PredictNothing,
					"--allowdialogs": complete.PredictNothing,
				},
			},
			"caches": complete.Command{
				Sub: complete.Commands{
					"bim360": complete.Command{
						Sub: complete.Commands{
							"clear": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{},
					},
				},
				Flags: complete.Flags{
					"--help": complete.PredictNothing,
				},
			},
			"config": complete.Command{
				Sub: complete.Commands{},
				Flags: complete.Flags{
					"--help": complete.PredictNothing,
					"--from": complete.PredictNothing,
				},
			},
			"configs": complete.Command{
				Sub: complete.Commands{
					"bincache": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"checkupdates": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"autoupdate": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"rocketmode": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"logs": complete.Command{
						Sub: complete.Commands{
							"none": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"verbose": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"debug": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"filelogging": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"startuptimeout": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"loadbeta": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"cpyversion": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"usercanupdate": complete.Command{
						Sub: complete.Commands{
							"yes": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"no": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"usercanextend": complete.Command{
						Sub: complete.Commands{
							"yes": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"no": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"usercanconfig": complete.Command{
						Sub: complete.Commands{
							"yes": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"no": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"colordocs": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"tooltipdebuginfo": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"routes": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"port": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"coreapi": complete.Command{
								Sub: complete.Commands{
									"enable": complete.Command{
										Sub: complete.Commands{},
										Flags: complete.Flags{
											"--log": complete.PredictNothing,
										},
									},
									"disable": complete.Command{
										Sub: complete.Commands{},
										Flags: complete.Flags{
											"--log": complete.PredictNothing,
										},
									},
								},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log":  complete.PredictNothing,
							"--help": complete.PredictNothing,
						},
					},
					"telemetry": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"utc": complete.Command{
								Sub: complete.Commands{
									"yes": complete.Command{
										Sub: complete.Commands{},
										Flags: complete.Flags{
											"--log": complete.PredictNothing,
										},
									},
									"no": complete.Command{
										Sub: complete.Commands{},
										Flags: complete.Flags{
											"--log": complete.PredictNothing,
										},
									},
								},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"file": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"server": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"hooks": complete.Command{
								Sub: complete.Commands{
									"yes": complete.Command{
										Sub: complete.Commands{},
										Flags: complete.Flags{
											"--log": complete.PredictNothing,
										},
									},
									"no": complete.Command{
										Sub: complete.Commands{},
										Flags: complete.Flags{
											"--log": complete.PredictNothing,
										},
									},
								},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log":  complete.PredictNothing,
							"--help": complete.PredictNothing,
						},
					},
					"apptelemetry": complete.Command{
						Sub: complete.Commands{
							"enable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"disable": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"flags": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
							"server": complete.Command{
								Sub: complete.Commands{},
								Flags: complete.Flags{
									"--log": complete.PredictNothing,
								},
							},
						},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"outputcss": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"seed": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--lock": complete.PredictNothing,
						},
					},
					"enable": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
					"disable": complete.Command{
						Sub: complete.Commands{},
						Flags: complete.Flags{
							"--log": complete.PredictNothing,
						},
					},
				},
				Flags: complete.Flags{
					"--log":  complete.PredictNothing,
					"--help": complete.PredictNothing,
				},
			},
			"doctor": complete.Command{
				Sub: complete.Commands{},
				Flags: complete.Flags{
					"--list":   complete.PredictNothing,
					"--help":   complete.PredictNothing,
					"--dryrun": complete.PredictNothing,
				},
			},
		},
		Flags: complete.Flags{
			"--debug":   complete.PredictNothing,
			"--verbose": complete.PredictNothing,
			"--help":    complete.PredictNothing,
			"--usage":   complete.PredictNothing,
			"--version": complete.PredictNothing,
		},
	}
	complete.New("pyrevit", pyrevit).Run()
}