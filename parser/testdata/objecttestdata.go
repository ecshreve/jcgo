package testdata

import "github.com/ecshreve/jcgo/parser"

type ObjectTestData struct {
	SimpleMapObj            *parser.MapObj
	SimpleNestedMapObj      *parser.MapObj
	DoubleNestedMapObj      *parser.MapObj
	SimpleSliceMapObj       *parser.MapObj
	SimpleSliceNestedMapObj *parser.MapObj
	ComplexMapObj           *parser.MapObj
}

func NewObjectTestData() *ObjectTestData {
	simpleMapObj := &parser.MapObj{
		Prefix: "",
		Val: []parser.Object{
			&parser.StringObj{
				Prefix: "key1",
				Val:    "val1",
			},
			&parser.StringObj{
				Prefix: "key2",
				Val:    "val2",
			},
			&parser.StringObj{
				Prefix: "key3",
				Val:    "val3",
			},
		},
	}

	simpleNestedMapObj := &parser.MapObj{
		Prefix: "",
		Val: []parser.Object{
			&parser.MapObj{
				Prefix: "outer1",
				Val: []parser.Object{
					&parser.StringObj{
						Prefix: "outer1_inner1",
						Val:    "innerval1",
					},
					&parser.StringObj{
						Prefix: "outer1_inner2",
						Val:    "innerval2",
					},
				},
			},
			&parser.StringObj{
				Prefix: "outer2",
				Val:    "outerval2",
			},
		},
	}

	doubleNestedMapObj := &parser.MapObj{
		Prefix: "",
		Val: []parser.Object{
			&parser.MapObj{
				Prefix: "outer1",
				Val: []parser.Object{
					&parser.StringObj{
						Prefix: "outer1_inner1",
						Val:    "innerval1",
					},
					&parser.MapObj{
						Prefix: "outer1_nestedmap",
						Val: []parser.Object{
							&parser.StringObj{
								Prefix: "outer1_nestedmap_nested1",
								Val:    "nestedval1",
							},
							&parser.StringObj{
								Prefix: "outer1_nestedmap_nested2",
								Val:    "nestedval2",
							},
						},
					},
				},
			},
			&parser.StringObj{
				Prefix: "outer2",
				Val:    "outerval2",
			},
		},
	}

	simpleSliceMapObj := &parser.MapObj{
		Prefix: "",
		Val: []parser.Object{
			&parser.SliceObj{
				Prefix: "data",
				Val: []parser.Object{
					&parser.MapObj{
						Prefix: "data",
						Val: []parser.Object{
							&parser.StringObj{
								Prefix: "data_key1",
								Val:    "val1",
							},
							&parser.StringObj{
								Prefix: "data_key2",
								Val:    "val2",
							},
							&parser.StringObj{
								Prefix: "data_key3",
								Val:    "val3",
							},
						},
					},
					&parser.MapObj{
						Prefix: "data",
						Val: []parser.Object{
							&parser.StringObj{
								Prefix: "data_key1",
								Val:    "val4",
							},
							&parser.StringObj{
								Prefix: "data_key2",
								Val:    "val5",
							},
							&parser.StringObj{
								Prefix: "data_key3",
								Val:    "val6",
							},
						},
					},
				},
			},
		},
	}

	simpleSliceNestedMapObj := &parser.MapObj{
		Prefix: "",
		Val: []parser.Object{
			&parser.SliceObj{
				Prefix: "data",
				Val: []parser.Object{
					&parser.MapObj{
						Prefix: "data",
						Val: []parser.Object{
							&parser.StringObj{
								Prefix: "data_key1",
								Val:    "val1",
							},
							&parser.MapObj{
								Prefix: "data_nestedmap",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_nestedmap_nested1",
										Val:    "nestedval1",
									},
									&parser.StringObj{
										Prefix: "data_nestedmap_nested2",
										Val:    "nestedval2",
									},
								},
							},
						},
					},
					&parser.MapObj{
						Prefix: "data",
						Val: []parser.Object{
							&parser.StringObj{
								Prefix: "data_key1",
								Val:    "val4",
							},
							&parser.MapObj{
								Prefix: "data_nestedmap",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_nestedmap_nested1",
										Val:    "nestedval3",
									},
									&parser.StringObj{
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

	complexMapObj := &parser.MapObj{
		Prefix: "",
		Val: []parser.Object{
			&parser.SliceObj{
				Prefix: "data",
				Val: []parser.Object{
					&parser.MapObj{
						Prefix: "data",
						Val: []parser.Object{
							&parser.StringObj{
								Prefix: "data_key1",
								Val:    "val1",
							},
							&parser.MapObj{
								Prefix: "data_nestedmap",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_nestedmap_nested1",
										Val:    "nestedval1",
									},
									&parser.StringObj{
										Prefix: "data_nestedmap_nested2",
										Val:    "nestedval2",
									},
								},
							},
							&parser.SliceObj{
								Prefix: "data_nestedslice",
								Val: []parser.Object{
									&parser.MapObj{
										Prefix: "data_nestedslice",
										Val: []parser.Object{
											&parser.StringObj{
												Prefix: "data_nestedslice_nestedslicemap1",
												Val:    "nestedslicemapval1",
											},
											&parser.StringObj{
												Prefix: "data_nestedslice_nestedslicemap2",
												Val:    "nestedslicemapval2",
											},
										},
									},
									&parser.MapObj{
										Prefix: "data_nestedslice",
										Val: []parser.Object{
											&parser.StringObj{
												Prefix: "data_nestedslice_nestedslicemap1",
												Val:    "nestedslicemapval3",
											},
											&parser.StringObj{
												Prefix: "data_nestedslice_nestedslicemap2",
												Val:    "nestedslicemapval4",
											},
										},
									},
								},
							},
						},
					},
					&parser.MapObj{
						Prefix: "data",
						Val: []parser.Object{
							&parser.StringObj{
								Prefix: "data_key1",
								Val:    "val4",
							},
							&parser.MapObj{
								Prefix: "data_nestedmap",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_nestedmap_nested1",
										Val:    "nestedval3",
									},
									&parser.StringObj{
										Prefix: "data_nestedmap_nested2",
										Val:    "nestedval4",
									},
								},
							},
							&parser.SliceObj{
								Prefix: "data_nestedslice",
								Val: []parser.Object{
									&parser.MapObj{
										Prefix: "data_nestedslice",
										Val: []parser.Object{
											&parser.StringObj{
												Prefix: "data_nestedslice_nestedslicemap1",
												Val:    "nestedslicemapval5",
											},
											&parser.StringObj{
												Prefix: "data_nestedslice_nestedslicemap2",
												Val:    "nestedslicemapval6",
											},
										},
									},
									&parser.MapObj{
										Prefix: "data_nestedslice",
										Val: []parser.Object{
											&parser.StringObj{
												Prefix: "data_nestedslice_nestedslicemap1",
												Val:    "nestedslicemapval7",
											},
											&parser.StringObj{
												Prefix: "data_nestedslice_nestedslicemap2",
												Val:    "nestedslicemapval8",
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
		SimpleNestedMapObj:      simpleNestedMapObj,
		DoubleNestedMapObj:      doubleNestedMapObj,
		SimpleSliceMapObj:       simpleSliceMapObj,
		SimpleSliceNestedMapObj: simpleSliceNestedMapObj,
		ComplexMapObj:           complexMapObj,
	}
}
