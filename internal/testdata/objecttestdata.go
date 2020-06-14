package testdata

import "github.com/ecshreve/jcgo/internal/object"

type ObjectTestData struct {
	SimpleMapObj            *object.MapObj
	SimpleAllTypesMapObj    *object.MapObj
	SimpleNestedMapObj      *object.MapObj
	DoubleNestedMapObj      *object.MapObj
	SimpleSliceMapObj       *object.MapObj
	SimpleSliceNestedMapObj *object.MapObj
	ComplexMapObj           *object.MapObj
}

func NewObjectTestData() *ObjectTestData {
	simpleMapObj := &object.MapObj{
		Prefix: "",
		Val: []object.Object{
			&object.StringObj{
				Prefix: "key1",
				Val:    "val1",
			},
			&object.StringObj{
				Prefix: "key2",
				Val:    "val2",
			},
			&object.StringObj{
				Prefix: "key3",
				Val:    "val3",
			},
		},
	}

	simpleAllTypesMapObj := &object.MapObj{
		Prefix: "",
		Val: []object.Object{
			&object.StringObj{
				Prefix: "key1",
				Val:    "val1",
			},
			&object.FloatObj{
				Prefix: "key2",
				Val:    float64(5),
			},
			&object.FloatObj{
				Prefix: "key3",
				Val:    float64(5.5),
			},
			&object.BoolObj{
				Prefix: "key4",
				Val:    true,
			},
		},
	}

	simpleNestedMapObj := &object.MapObj{
		Prefix: "",
		Val: []object.Object{
			&object.MapObj{
				Prefix: "outer1",
				Val: []object.Object{
					&object.StringObj{
						Prefix: "outer1_inner1",
						Val:    "innerval1",
					},
					&object.StringObj{
						Prefix: "outer1_inner2",
						Val:    "innerval2",
					},
				},
			},
			&object.StringObj{
				Prefix: "outer2",
				Val:    "outerval2",
			},
		},
	}

	doubleNestedMapObj := &object.MapObj{
		Prefix: "",
		Val: []object.Object{
			&object.MapObj{
				Prefix: "outer1",
				Val: []object.Object{
					&object.StringObj{
						Prefix: "outer1_inner1",
						Val:    "innerval1",
					},
					&object.MapObj{
						Prefix: "outer1_nestedmap",
						Val: []object.Object{
							&object.StringObj{
								Prefix: "outer1_nestedmap_nested1",
								Val:    "nestedval1",
							},
							&object.StringObj{
								Prefix: "outer1_nestedmap_nested2",
								Val:    "nestedval2",
							},
						},
					},
				},
			},
			&object.StringObj{
				Prefix: "outer2",
				Val:    "outerval2",
			},
		},
	}

	simpleSliceMapObj := &object.MapObj{
		Prefix: "",
		Val: []object.Object{
			&object.SliceObj{
				Prefix: "data",
				Val: []object.Object{
					&object.MapObj{
						Prefix: "data",
						Val: []object.Object{
							&object.StringObj{
								Prefix: "data_key1",
								Val:    "val1",
							},
							&object.StringObj{
								Prefix: "data_key2",
								Val:    "val2",
							},
							&object.StringObj{
								Prefix: "data_key3",
								Val:    "val3",
							},
						},
					},
					&object.MapObj{
						Prefix: "data",
						Val: []object.Object{
							&object.StringObj{
								Prefix: "data_key1",
								Val:    "val4",
							},
							&object.StringObj{
								Prefix: "data_key2",
								Val:    "val5",
							},
							&object.StringObj{
								Prefix: "data_key3",
								Val:    "val6",
							},
						},
					},
				},
			},
		},
	}

	simpleSliceNestedMapObj := &object.MapObj{
		Prefix: "",
		Val: []object.Object{
			&object.SliceObj{
				Prefix: "data",
				Val: []object.Object{
					&object.MapObj{
						Prefix: "data",
						Val: []object.Object{
							&object.StringObj{
								Prefix: "data_key1",
								Val:    "val1",
							},
							&object.MapObj{
								Prefix: "data_nestedmap",
								Val: []object.Object{
									&object.StringObj{
										Prefix: "data_nestedmap_nested1",
										Val:    "nestedval1",
									},
									&object.StringObj{
										Prefix: "data_nestedmap_nested2",
										Val:    "nestedval2",
									},
								},
							},
						},
					},
					&object.MapObj{
						Prefix: "data",
						Val: []object.Object{
							&object.StringObj{
								Prefix: "data_key1",
								Val:    "val4",
							},
							&object.MapObj{
								Prefix: "data_nestedmap",
								Val: []object.Object{
									&object.StringObj{
										Prefix: "data_nestedmap_nested1",
										Val:    "nestedval3",
									},
									&object.StringObj{
										Prefix: "data_nestedmap_nested2",
										Val:    "nestedval4",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	complexMapObj := &object.MapObj{
		Prefix: "",
		Val: []object.Object{
			&object.SliceObj{
				Prefix: "data",
				Val: []object.Object{
					&object.MapObj{
						Prefix: "data",
						Val: []object.Object{
							&object.StringObj{
								Prefix: "data_key1",
								Val:    "val1",
							},
							&object.MapObj{
								Prefix: "data_nestedmap",
								Val: []object.Object{
									&object.StringObj{
										Prefix: "data_nestedmap_nested1",
										Val:    "nestedval1",
									},
									&object.StringObj{
										Prefix: "data_nestedmap_nested2",
										Val:    "nestedval2",
									},
								},
							},
							&object.SliceObj{
								Prefix: "data_nestedslice",
								Val: []object.Object{
									&object.MapObj{
										Prefix: "data_nestedslice",
										Val: []object.Object{
											&object.SliceObj{
												Prefix: "data_nestedslice_after",
												Val: []object.Object{
													&object.MapObj{
														Prefix: "data_nestedslice_after",
														Val: []object.Object{
															&object.StringObj{
																Prefix: "data_nestedslice_after_nestedslicemap1",
																Val:    "nestedslicemapval5",
															},
															&object.StringObj{
																Prefix: "data_nestedslice_after_nestedslicemap2",
																Val:    "nestedslicemapval6",
															},
														},
													},
													&object.MapObj{
														Prefix: "data_nestedslice_after",
														Val: []object.Object{
															&object.StringObj{
																Prefix: "data_nestedslice_after_nestedslicemap1",
																Val:    "nestedslicemapval7",
															},
															&object.StringObj{
																Prefix: "data_nestedslice_after_nestedslicemap2",
																Val:    "nestedslicemapval8",
															},
														},
													},
												},
											},
											&object.SliceObj{
												Prefix: "data_nestedslice_before",
												Val: []object.Object{
													&object.MapObj{
														Prefix: "data_nestedslice_before",
														Val: []object.Object{
															&object.StringObj{
																Prefix: "data_nestedslice_before_nestedslicemap1",
																Val:    "nestedslicemapval1",
															},
															&object.StringObj{
																Prefix: "data_nestedslice_before_nestedslicemap2",
																Val:    "nestedslicemapval2",
															},
														},
													},
													&object.MapObj{
														Prefix: "data_nestedslice_before",
														Val: []object.Object{
															&object.StringObj{
																Prefix: "data_nestedslice_before_nestedslicemap1",
																Val:    "nestedslicemapval3",
															},
															&object.StringObj{
																Prefix: "data_nestedslice_before_nestedslicemap2",
																Val:    "nestedslicemapval4",
															},
														},
													},
												},
											},
										},
									},
									&object.MapObj{
										Prefix: "data_nestedslice",
										Val: []object.Object{
											&object.SliceObj{
												Prefix: "data_nestedslice_after",
												Val: []object.Object{
													&object.MapObj{
														Prefix: "data_nestedslice_after",
														Val: []object.Object{
															&object.StringObj{
																Prefix: "data_nestedslice_after_nestedslicemap1",
																Val:    "nestedslicemapval15",
															},
															&object.StringObj{
																Prefix: "data_nestedslice_after_nestedslicemap2",
																Val:    "nestedslicemapval16",
															},
														},
													},
													&object.MapObj{
														Prefix: "data_nestedslice_after",
														Val: []object.Object{
															&object.StringObj{
																Prefix: "data_nestedslice_after_nestedslicemap1",
																Val:    "nestedslicemapval17",
															},
															&object.StringObj{
																Prefix: "data_nestedslice_after_nestedslicemap2",
																Val:    "nestedslicemapval18",
															},
														},
													},
												},
											},
											&object.SliceObj{
												Prefix: "data_nestedslice_before",
												Val: []object.Object{
													&object.MapObj{
														Prefix: "data_nestedslice_before",
														Val: []object.Object{
															&object.StringObj{
																Prefix: "data_nestedslice_before_nestedslicemap1",
																Val:    "nestedslicemapval11",
															},
															&object.StringObj{
																Prefix: "data_nestedslice_before_nestedslicemap2",
																Val:    "nestedslicemapval12",
															},
														},
													},
													&object.MapObj{
														Prefix: "data_nestedslice_before",
														Val: []object.Object{
															&object.StringObj{
																Prefix: "data_nestedslice_before_nestedslicemap1",
																Val:    "nestedslicemapval13",
															},
															&object.StringObj{
																Prefix: "data_nestedslice_before_nestedslicemap2",
																Val:    "nestedslicemapval14",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return &ObjectTestData{
		SimpleMapObj:            simpleMapObj,
		SimpleAllTypesMapObj:    simpleAllTypesMapObj,
		SimpleNestedMapObj:      simpleNestedMapObj,
		DoubleNestedMapObj:      doubleNestedMapObj,
		SimpleSliceMapObj:       simpleSliceMapObj,
		SimpleSliceNestedMapObj: simpleSliceNestedMapObj,
		ComplexMapObj:           complexMapObj,
	}
}
